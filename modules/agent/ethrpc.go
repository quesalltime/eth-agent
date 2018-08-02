package agent

import (
	"bytes"
	"encoding/json"
	"eth-agent/config"
	"eth-agent/modules/agent/struct/rqst"
	"fmt"

	"eth-agent/modules/logger"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

var (
	addressLength42 = 42
	addressLength66 = 66
)

// JSONResponse : response from the JSONRPC call
// type JSONResponse struct {
// 	Jsonrpc string `json:"jsonrpc"`
// 	Result  string `json:"result"`
// 	ID      int    `json:"id"`
// }

// ParseCookie parse the SetCookie string from http response in array format [0]: AWSALB, [1]: Expires, [2]: Path
func ParseCookie(jsonRPCResponse *http.Response) Cookie {
	setCookie := jsonRPCResponse.Header.Get("Set-Cookie")

	SpilitCookies := strings.Split(setCookie, ";")
	cookiesInforamtion := Cookie{
		AWSALB:  (SpilitCookies[0])[7:],
		Expires: (SpilitCookies[1])[9:],
		Path:    (SpilitCookies[2])[6:],
	}

	return cookiesInforamtion
}

// ParseJSONRPCResponse print the json rpc result
func ParseJSONRPCResponse(jsonRPCResponse *http.Response) interface{} {

	resBody, err := ioutil.ReadAll(jsonRPCResponse.Body)
	if err != nil {
		logger.Console().Panic(err)
		logger.File().Panic(err)
	}

	var jResponse interface{}
	err = json.Unmarshal(resBody, &jResponse)

	if err != nil {
		logger.Console().Panic(err)
		logger.File().Panic(err)
	}
	// defer jsonRPCResponse.Body.Close()
	return jResponse
}

// requestNative create json-RPC method call along with AWSALB cookie
func requestNative(payload rqst.Payload, cookie string, expires int) (interface{}, Cookie) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		logger.Console().Panic(err)
		logger.File().Panic(err)
	}

	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest(http.MethodPost, config.EthURL, body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")

	if cookie != "" {
		req.AddCookie(&http.Cookie{
			Name:    "AWSALB",
			Value:   cookie,
			Expires: time.Now().Add(time.Duration(expires) * time.Second),
		})
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Console().Warn(err)
		logger.File().Warn(err)
	}
	return ParseJSONRPCResponse(resp), ParseCookie(resp)
}

// GetRedisInstance get redis client in singleton pattern
func GetRedisInstance(strPasswd string, strIP string, strPort string, strDB int) (*redis.Client, error) {
	var err error
	client := redis.NewClient(&redis.Options{
		Addr:     strIP + ":" + strPort,
		Password: strPasswd, // no password set
		DB:       strDB,     // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong, err)
	return client, err
}

// ReleaseRedisInstance close Redis instance
func ReleaseRedisInstance(db *redis.Client) {
	db.Close()
}
