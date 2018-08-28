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

func GetCode(data rqst.Payload) interface{} {
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
	contractAddress := param.(string)
	contractAddress, err = historyUtilsCommon.Check40HeximalFormat(contractAddress)
	if err != nil {
		errors := common.Error{
			ErrorType:        0,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
	}

	fmt.Println("================getCode initial parameter================")
	fmt.Printf("contractAddress:%s\n", contractAddress)

	var response rsps.GetCodeResponse

	response = getCodeIndexer(contractAddress)
	return response
}

func getCodeIndexer(contractAddress string) rsps.GetCodeResponse {
	var response rsps.GetCodeResponse
	response.Jsonrpc = "2.0"
	response.ID = 73

	condition := bson.M{
		"contractAddress": contractAddress,
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
