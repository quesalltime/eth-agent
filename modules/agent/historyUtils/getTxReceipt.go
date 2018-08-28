package historyUtils

import (
	"errors"
	"eth-agent/common"
	historyUtilsCommon "eth-agent/modules/agent/historyUtils/common"
	model "eth-agent/modules/agent/historyUtils/model"
	"eth-agent/modules/agent/historyUtils/struct/rsps"
	"fmt"

	"eth-agent/modules/agent/struct/rqst"
	"eth-agent/modules/logger"

	"gopkg.in/mgo.v2/bson"
)

func GetTxReceipt(data rqst.Payload) interface{} {
	var err error
	var message string

	params := data.Params
	if len(params) != 1 {
		message = fmt.Sprintf("Invalid params: invalid length 2, expected 1 elements in array.")
		err = errors.New(message)
		errors := common.Error{
			ErrorType:        0,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
	}
	param := params[0]
	transactionAddress := param.(string)
	transactionAddress, err = historyUtilsCommon.Check64HeximalFormat(transactionAddress)
	if err != nil {
		errors := common.Error{
			ErrorType:        0,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
	}

	fmt.Println("================getTransactionReceipt initial parameter================")
	fmt.Printf("transactionReceipt:%s\n", transactionAddress)

	var response rsps.ReceiptResponse

	response = getTxReceiptIndexer(transactionAddress)
	return response
}

func getTxReceiptIndexer(transactionAddress string) rsps.ReceiptResponse {
	var response rsps.ReceiptResponse
	response.Jsonrpc = "2.0"
	response.ID = 73

	condition := bson.M{
		"transactionHash": transactionAddress,
	}

	result, err := model.RetrieveReceipts(condition)

	if err != nil {
		errors := common.Error{
			ErrorType:        1,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
		logger.File().Error(err)
	}

	// Convert the blockNumber from decimal to hex...
	// The blockNumber in BS_Receipt is heximal format (string)
	// result[0].BlockNumber = historyUtilsCommon.ParseInt64ToHex(result[0].BlockNumber.(int))

	response.Result = result[0]
	return response
}
