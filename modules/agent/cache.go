package agent

import (
	"encoding/json"
	"eth-agent/config"
	"eth-agent/modules/agent/struct/rqst"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func requestCache(data rqst.Payload, userAPIKey, AWSALBCookie string) (interface{}, Cookie) {
	var cookie Cookie
	// Redis instance
	redisUser, _ := GetRedisInstance(config.SysConf.Redis.Password, config.SysConf.Redis.Domain, config.SysConf.Redis.Port, 0)
	// If there is no user api key store in redis.
	if AWSALBCookie == "" {
		fmt.Printf("Userkey %s does not exist \n", userAPIKey)
		// Create a json rpc call to get new AWSALB cookie...
		jsonRPCResponse, AWSCookie := requestNative(data, "", JSONRPCMethodCookieExipres)
		cookie = AWSCookie
		// // Seting AWSALB
		// reddisErr := redisUser.Set(userAPIKey, cookie.AWSALB, time.Duration(expiresForAWSALBTime)*time.Second).Err()
		// if reddisErr != nil {
		// 	panic("redisUser set error")
		// } else {
		// 	fmt.Printf("Set userKey:%s, AWSALB  cookie: %s \n", userAPIKey, cookie.AWSALB)
		// }
		return jsonRPCResponse, cookie
	}

	fmt.Println("userkey's AWSALB:", AWSALBCookie)
	userAPIWithMethod := userAPIKey + ":" + data.Method
	userAPIWithMethodResponseInRedis, userAPIWithMethodResponseInRedisErr := redisUser.Get(userAPIWithMethod).Result()
	// If there's no data in redis, then asking parity for response and doing cache...
	if userAPIWithMethodResponseInRedisErr == redis.Nil {
		JSONRPCResponseJSONFormat, _ := requestNative(data, AWSALBCookie, JSONRPCMethodCookieExipres)
		JSONRPCResponseStringFormat, jsonMarshalErr := json.Marshal(JSONRPCResponseJSONFormat)

		if jsonMarshalErr != nil {
			panic("Error Message: serAPIWithJSONmethod error")
		}

		// No json method cache in redis. do saving...
		setUserAPIAndResultError := redisUser.Set(userAPIWithMethod, JSONRPCResponseStringFormat, time.Duration(JSONRPCMethodCookieExipres)*time.Second).Err()
		if setUserAPIAndResultError != nil {
			panic("Error Message: setUserAPIAndResult error")
		}

		fmt.Printf("Set %s. Cache data: %s \n", userAPIWithMethod, JSONRPCResponseStringFormat)
		fmt.Printf("Cached data for %d seconds... \n", JSONRPCMethodCookieExipres)
		return JSONRPCResponseJSONFormat, cookie

	}

	fmt.Printf("%s is exist: %s. \n", userAPIWithMethod, userAPIWithMethodResponseInRedis)
	var JSONRPCMethodResponse interface{}
	jsonUnmarshalErr := json.Unmarshal([]byte(userAPIWithMethodResponseInRedis), &JSONRPCMethodResponse)

	if jsonUnmarshalErr != nil {
		panic(jsonUnmarshalErr)
	}
	return JSONRPCMethodResponse, cookie
}
