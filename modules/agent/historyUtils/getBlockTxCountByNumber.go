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

// GetBlockTxCountByNumber handles counting the total amount of transaction in a block by giving it's block height.
func GetBlockTxCountByNumber(data rqst.Payload) interface{} {
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
	blockNumber := param.(string)

	err = historyUtilsCommon.CheckBlockNumberFormat(blockNumber)
	if err != nil {
		errors := common.Error{
			ErrorType:        0,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
	}

	blockNumberInteger, err := historyUtilsCommon.ParseHexToInt64(blockNumber)
	if err != nil {
		errors := common.Error{
			ErrorType:        0,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
	}
	fmt.Println("================getBlockTransactionCountByNumber initial parameter================")
	fmt.Printf("blockNumber:%s\n", blockNumber)

	var response rsps.GBTCResponse

	response = getBlockTxCountByNumber(blockNumberInteger)
	return response
}

func getBlockTxCountByNumber(blockNumber int64) rsps.GBTCResponse {
	var response rsps.GBTCResponse
	response.Jsonrpc = "2.0"
	response.ID = 73

	condition := bson.M{
		"number": blockNumber,
	}

	result, err := model.RetrieveBlock(condition)
	if err != nil {
		errors := common.Error{
			ErrorType:        1,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
		logger.File().Error(err)
	}
	// check the blockHeigh is exsit or not.
	// if not exist, then the result will be nil (lenght == 0)
	if len(result) == 0 {
		response.Result = "0x0"
	} else {
		lengthOfTx := len(result[0].Transactions)
		response.Result = historyUtilsCommon.ParseInt64ToHex(lengthOfTx)
	}
	return response
}
