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
		response = historyUtils.GetBlockTxCountByNumber(data)
	case "eth_getBlockTransactionCountByHash":
		response = historyUtils.GetBlockTxCountByHash(data)
	case "eth_getTransactionReceipt":
		response = historyUtils.GetTxReceipt(data)
	case "eth_getCode":
		response = historyUtils.GetCode(data)
	case "eth_getBlockByHash":
		response = historyUtils.GetBlockByHash(data)
	case "eth_getBlockByNumber":
		response = historyUtils.GetBlockByNumber(data)
	default:
		response = "no corresponding method"
		logger.Console().Debug(fmt.Sprintf("%s", response))
	}
	return response
}

func notImplemented() string {
	return "not implemented yet"
}
