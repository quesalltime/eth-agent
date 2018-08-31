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

// GetCode will return the code of contrac by giving two parameters:(transactino hash, block height)
func GetCode(data rqst.Payload) interface{} {
	var err error
	var message string

	params := data.Params
	if len(params) != 2 {
		// [contract address] [block number、"latest"、"earlest"、"pending"]
		// We temporary forbid user to put "latest","earlest","pending" in parameters.
		message = fmt.Sprintf("missing value for required argument.")
		err = errors.New(message)
		errors := common.Error{
			ErrorType:        0,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
	}

	contractAddress := params[0].(string)
	contractAddress, err = historyUtilsCommon.Check40HeximalFormat(contractAddress)
	if err != nil {
		errors := common.Error{
			ErrorType:        0,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
	}

	blockNumber := params[1].(string)
	err = historyUtilsCommon.CheckBlockNumberFormat(blockNumber)
	if err != nil {
		errors := common.Error{
			ErrorType:        0,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
	}

	fmt.Println("================getCode initial parameter================")
	fmt.Printf("contractAddress:%s, blockNumber:%s \n", contractAddress, blockNumber)

	var response rsps.GetCodeResponse

	response = getCodeIndexer(contractAddress, blockNumber)
	return response
}

func getCodeIndexer(contractAddress string, blockNumber string) rsps.GetCodeResponse {
	var response rsps.GetCodeResponse
	response.Jsonrpc = "2.0"
	response.ID = 73

	condition := bson.M{
		"contractAddress": contractAddress,
		"blockNumber":     blockNumber,
	}

	result, err := model.RetrieveContracts(condition)
	if err != nil {
		errors := common.Error{
			ErrorType:        1,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
		logger.File().Error(err)
	}

	response.Result = result[0].Code
	return response
}
