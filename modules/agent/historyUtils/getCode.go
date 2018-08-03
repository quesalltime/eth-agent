package historyUtils

import (
	"errors"
	"eth-agent/common"
	historyUtilsCommon "eth-agent/modules/agent/historyUtils/common"
	historyUtilsMongo "eth-agent/modules/agent/historyUtils/mongo"
	contractStruct "eth-agent/modules/agent/historyUtils/struct/bs_contract"
	"eth-agent/modules/agent/historyUtils/struct/rsps"
	"fmt"

	"eth-agent/modules/agent/struct/rqst"
	"eth-agent/modules/logger"

	"gopkg.in/mgo.v2/bson"
)

var (
	BS_Contracts = "BS_Contracts"
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

	response = GetCodeIndexer(contractAddress)
	return response
}

func GetCodeIndexer(contractAddress string) rsps.GetCodeResponse {
	var response rsps.GetCodeResponse
	response.Jsonrpc = "2.0"
	response.ID = 73

	condition := bson.M{
		"contractAddress": contractAddress,
	}

	result, err := RetrieveContract(condition)
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

// RetrieveBlock retrieve specific block data from mongo
func RetrieveContract(conditions map[string]interface{}) ([]contractStruct.Contract, error) {
	var err error

	mongo, err := historyUtilsMongo.GetMongoSession()
	if err != nil {
		errors := common.Error{
			ErrorType:        1,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
		logger.File().Error(err)
	}

	defer mongo.Close()

	collection := mongo.DB(dbName).C(BS_Contracts)
	result := []contractStruct.Contract{}
	err = collection.Find(conditions).All(&result)

	if err != nil {
		message := fmt.Sprintf("Retrive Receipt of Transaction failded")
		err = errors.New(message)
	}

	return result, err
}