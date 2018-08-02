package agent

import (
	"eth-agent/config"
	"eth-agent/modules/agent/struct/rqst"
	"eth-agent/modules/logger"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	// ethURL     string
	httpClient = &http.Client{}
)

var (
	AWSALBCookieExpires        = config.AWSALBCookieExpires
	JSONRPCMethodCookieExipres = config.JSONRPCMethodCookieExipres
)

var rpcTypeMap = map[string]int{
	// to Parity
	"eth_estimateGas":                   nativeType,
	"eth_sendTransaction":               nativeType,
	"eth_sendRawTransaction":            nativeType,
	"eth_newBlockFilter":                nativeType,
	"eth_newFilter":                     nativeType,
	"eth_newPendingTransactionFilter":   nativeType,
	"eth_getFilterChanges":              nativeType,
	"eth_getFilterLogs":                 nativeType,
	"eth_protocolVersion":               nativeType,
	"eth_getUncleCountByBlockHash":      nativeType,
	"eth_getUncleCountByBlockNumber":    nativeType,
	"eth_getUncleByBlockHashAndIndex":   nativeType,
	"eth_getUncleByBlockNumberAndIndex": nativeType,
	"eth_blockNumber":                   nativeType,
	"eth_getTransactionCount":           nativeType,

	// to Memeory-cache
	"eth_getBalance": cachedType,
	"eth_call":       cachedType,

	// to MongoDB
	"eth_getBlockTransactionCountByNumber": historyType,
	"eth_getBlockTransactionCountByHash":   historyType,
	"eth_getLogs":                          historyType,
	"eth_getTransactionReceipt":            historyType,
	"eth_getCode":                          historyType,
	"eth_getBlockByHash":                   historyType,
	"eth_getBlockByNumber":                 historyType,
}

const (
	nativeType  = 0
	cachedType  = 1
	historyType = 2
)

// Cookie : information from AWS loadbalancer
type Cookie struct {
	AWSALB  string
	Expires string
	Path    string
}

// Redirect - redirect request from bm router
func Redirect(c *gin.Context) {
	userAPIKey := c.Request.Header.Get("Authorization")
	reqBody := map[string]interface{}{}
	err := c.ShouldBindJSON(&reqBody)
	if err != nil {
		logger.Console().Panic(err)
		logger.File().Panic(err)
		return
	}

	result := redirectRPC(reqBody, userAPIKey)
	c.JSON(http.StatusOK, result)

	// rawData, err := c.GetRawData()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// req, err := http.NewRequest(http.MethodPost, config.EthURL, bytes.NewBuffer(rawData))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// req.Header.Set("Content-Type", "application/json")

	// resp, err := httpClient.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer resp.Body.Close()

	// body, _ := ioutil.ReadAll(resp.Body)
	// var arbitraryObj interface{}
	// json.Unmarshal(body, &arbitraryObj)
	// c.JSON(resp.StatusCode, arbitraryObj)
}

func getAWSALBCookie(userAPIKey string) (Cookie, error) {
	// Redis instance
	redis, _ := GetRedisInstance(config.SysConf.Redis.Password, config.SysConf.Redis.Domain, config.SysConf.Redis.Port, 0)
	// Setting user api key.
	AWSALBCookie, err := redis.Get(userAPIKey).Result()
	var cookie Cookie
	if err != nil {
		logger.Console().Warn(err)
		logger.File().Warn(err)
		return cookie, err
	}

	cookie = Cookie{
		AWSALB:  AWSALBCookie,
		Expires: "",
		Path:    "",
	}
	return cookie, nil

}

func setAWSALBCookie(cookie Cookie, userAPIKey string) bool {
	redis, _ := GetRedisInstance(config.SysConf.Redis.Password, config.SysConf.Redis.Domain, config.SysConf.Redis.Port, 0)
	// Seting AWSALB
	err := redis.Set(userAPIKey, cookie.AWSALB, time.Duration(AWSALBCookieExpires)*time.Second).Err()
	if err != nil {
		logger.Console().Warn(err)
		logger.File().Warn(err)
		return false
	}

	logger.Console().Debug(fmt.Sprintf("Set userKey:%s, AWSALB  cookie: %s \n", userAPIKey, cookie.AWSALB))
	return true
}

func explainType(params []interface{}) []interface{} {
	paramsWithType := make([]interface{}, reflect.ValueOf(params).Len())

	for i, value := range params {
		switch target := value.(type) {
		case int:
			paramsWithType[i] = params[i].(int)
		case float64:
			paramsWithType[i] = params[i].(float64)
		case string:
			paramsWithType[i] = params[i].(string)
		case bool:
			paramsWithType[i] = params[i].(bool)
		case []interface{}:
			paramsWithType[i] = params[i].([]interface{})
		case map[string]interface{}:
			paramsWithType[i] = params[i].(map[string]interface{})
		default:
			logger.Console().Panic(fmt.Sprintf("The type of the parameter: %v ", target))
		}
	}
	return paramsWithType
}

func redirectRPC(reqBody map[string]interface{}, userAPIKey string) interface{} {
	method := reqBody["method"].(string)
	params := reqBody["params"].([]interface{})
	typedParams := explainType(params)

	// Cache and DB both need userAPIKey information
	// AWSALB cookie expired for about 1600
	oldCookie, _ := getAWSALBCookie(userAPIKey)

	payload := rqst.Payload{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  typedParams,
		ID:      64,
	}

	var result interface{}
	var cookie Cookie

	switch rpcType := rpcTypeMap[method]; rpcType {
	case nativeType:
		result, cookie = requestNative(payload, oldCookie.AWSALB, JSONRPCMethodCookieExipres)
		CheckCookieLiveness(cookie, oldCookie, userAPIKey)
	case cachedType:
		result, cookie = requestCache(payload, userAPIKey, oldCookie.AWSALB)
		CheckCookieLiveness(cookie, oldCookie, userAPIKey)
	case historyType:
		result = requestDB(payload)
	default:
		result, cookie = requestNative(payload, oldCookie.AWSALB, JSONRPCMethodCookieExipres)
	}

	return result
}

// CheckCookieLiveness check oldcookie is expired or not, if oldCookie.AWSALB is empty, then set a new one in redis.
func CheckCookieLiveness(newCookie Cookie, oldCookie Cookie, userAPIKey string) {

	if oldCookie.AWSALB == "" {
		if setAWSALBCookie(newCookie, userAPIKey) {
			logger.Console().Debug("Set new user APIKey success")
		}
	}
}
