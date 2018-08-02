package agent

import (
	"eth-agent/modules/agent/historyUtils"
	"eth-agent/modules/agent/struct/rqst"
	"fmt"

	"eth-agent/modules/logger"
)

func requestDB(data rqst.Payload) interface{} {

	payloadMethod := data.Method
	var response interface{}

	switch payloadMethod {
	case "eth_getLogs":
		response = historyUtils.GetLogs(data)
	case "eth_getBlockTransactionCountByNumber":
		response = historyUtils.GetBlockTransactionCountByNumber(data)
	case "eth_getBlockTransactionCountByHash":
		response = notImplemented()
	case "eth_getTransactionReceipt":
		response = notImplemented()
	case "eth_getCode":
		response = notImplemented()
	case "eth_getBlockByHash":
		response = notImplemented()
	case "eth_getBlockByNumber":
		response = notImplemented()
	default:
		response = "no corresponding method"
		logger.Console().Debug(fmt.Sprintf("%s", response))
	}
	return response
}

func notImplemented() string {
	return "not implemented yet"
}
