package historyUtils

import (
	"errors"
	"eth-agent/common"
	collectionName "eth-agent/modules/agent/historyUtils/common"
	historyUtilsCommon "eth-agent/modules/agent/historyUtils/common"
	historyUtilsMongo "eth-agent/modules/agent/historyUtils/mongo"
	receiptStruct "eth-agent/modules/agent/historyUtils/struct/bs_receipt"
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

	var response rsps.ReceiptReponse

	response = GetTxReceiptIndexer(transactionAddress)
	return response
}

func GetTxReceiptIndexer(transactionAddress string) rsps.ReceiptReponse {
	var response rsps.ReceiptReponse
	response.Jsonrpc = "2.0"
	response.ID = 73

	condition := bson.M{
		"transactionHash": transactionAddress,
	}

	result, err := RetrieveReceipt(condition)
	if err != nil {
		errors := common.Error{
			ErrorType:        1,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
		logger.File().Error(err)
	}

	// Convert the blockNumber from decimal to hex...

	result[0].BlockNumber = historyUtilsCommon.ParseInt64ToHex(result[0].BlockNumber.(int))

	response.Result = result[0]
	return response
}

// RetrieveBlock retrieve specific block data from mongo
func RetrieveReceipt(conditions map[string]interface{}) ([]receiptStruct.Receipt, error) {
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

	collection := mongo.DB(dbName).C(collectionName.BsReceipts)
	result := []receiptStruct.Receipt{}
	err = collection.Find(conditions).All(&result)

	if err != nil {
		message := fmt.Sprintf("Retrive Receipt of Transaction failded")
		err = errors.New(message)
	}

	return result, err
}
