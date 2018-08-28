package historyUtils

import (
	"errors"
	"eth-agent/common"
	"eth-agent/config"
	collectionName "eth-agent/modules/agent/historyUtils/common"
	historyUtilsCommon "eth-agent/modules/agent/historyUtils/common"
	historyUtilsMongo "eth-agent/modules/agent/historyUtils/mongo"
	receiptStrcut "eth-agent/modules/agent/historyUtils/struct/bs_receipt"
	"eth-agent/modules/agent/historyUtils/struct/rsps"
	"eth-agent/modules/agent/struct/rqst"
	"eth-agent/modules/logger"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

var (
	dbName = config.SysConf.Mongo.DBName
	// limit the number of the logs user can request.
	numberOfLimitation = (float64)(1000)
)

// GetLogs will check every parameters first.
// If there is blockhash in parameter,the parameter:fromBlock, toBlock is not allowed.
// Then, parse the result from mongo and return it to the request.
func GetLogs(data rqst.Payload) interface{} {
	var message string
	var err error

	params := data.Params
	// Ｃheck whether if the len of params is == 1
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
	// Need do loop again if there's type of value is []interface{} in array.
	paramObject := param.(map[string]interface{})
	// (1) If no fromBlock, toBlock ,then default = "latest"
	// (2) fromBlock, tolock, address should be examined whether it has "0x" prefix
	// (3) check fromBlock or toBlock is null or not.
	// Check fromBlock parameter
	paramFromBlock := paramObject["fromBlock"]
	fromBlockNumberIndex, err := historyUtilsCommon.CheckBlockNumberIndexFormat(paramFromBlock)
	if err != nil {
		errors := common.Error{
			ErrorType:        0,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
	}
	// Check the parameter: fromBlock
	paramObject["fromBlock"] = fromBlockNumberIndex
	paramToBlock := paramObject["toBlock"]
	toBlockNumberIndex, err := historyUtilsCommon.CheckBlockNumberIndexFormat(paramToBlock)
	if err != nil {
		errors := common.Error{
			ErrorType:        0,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
	}
	// Check the parameter: toBlock
	paramObject["toBlock"] = toBlockNumberIndex
	paramAddress := paramObject["address"]
	addressCheck, err := historyUtilsCommon.Check40HeximalFormat(paramAddress)
	if err != nil {
		errors := common.Error{
			ErrorType:        0,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
	}
	paramObject["address"] = addressCheck
	/* It is hard to withdraw those section above to ethrpc.go */

	// Check the parameter: topics
	switch paramObject["topics"].(type) {
	case nil:
		paramObject["topics"] = nil
	case []interface{}:
		topics := paramObject["topics"]
		topicsArray := topics.([]interface{})
		// Initial string array: []string, using for loop to initialize each interface typed var in interface[]
		for i, topic := range topicsArray {
			switch topic.(type) {
			case string:
				topicString := topic.(string)
				if topicString[0:2] != "0x" {
					message = fmt.Sprintf("invalid argument address: hex string without 0x prefix")
					err = errors.New(message)
					errors := common.Error{
						ErrorType:        0,
						ErrorDescription: err.Error(),
					}
					logger.Console().Panic(errors)
				}
				if len(topicString) != historyUtilsCommon.AddressLength66 {
					message = fmt.Sprintf("Invalid argument: topic[%d], hex has invalid length %d, should be %d (include 0x prefix)", i, len(topicString), historyUtilsCommon.AddressLength66)
					err = errors.New(message)
					errors := common.Error{
						ErrorType:        0,
						ErrorDescription: err.Error(),
					}
					logger.Console().Panic(errors)
				}
			default:
				message = fmt.Sprintf("The type of topic[%d] is incorrect, should be string", i)
				err = errors.New(message)
				errors := common.Error{
					ErrorType:        0,
					ErrorDescription: err.Error(),
				}
				logger.Console().Panic(errors)
			}
		}
	default:
		message = fmt.Sprintf("Address type should be string format with 0x prefix")
		err = errors.New(message)
		errors := common.Error{
			ErrorType:        0,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
		logger.File().Error(errors)

		return err
	}
	// Check the parameter: limit
	switch paramObject["limit"].(type) {
	case nil:
		paramObject["limit"] = (float64)(0)
	case float64:
		limit := paramObject["limit"]
		limitfloat64 := limit.(float64)
		if limitfloat64 > numberOfLimitation {
			message = fmt.Sprintf("Number of limit is too large, should be less than %d", numberOfLimitation)
			err = errors.New(message)
			errors := common.Error{
				ErrorType:        0,
				ErrorDescription: err.Error(),
			}
			logger.Console().Panic(errors)
			logger.File().Error(errors)
		}
	default:
		message := fmt.Sprintf("Type of Limit is not integer")
		err = errors.New(message)
		errors := common.Error{
			ErrorType:        0,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
		logger.File().Error(errors)
	}

	fromBlock := paramObject["fromBlock"].(string)
	toBlock := paramObject["toBlock"].(string)
	address := paramObject["address"].(string)
	topics := paramObject["topics"].([]interface{})
	limit := (int)(paramObject["limit"].(float64))

	fmt.Println("================getLogs initial parameter================")
	fmt.Printf("fromBlock:%s, toBlock:%s, address: %s, topics:%s, limit:%d\n", fromBlock, toBlock, address, topics, limit)

	var response rsps.GetLogsResponse

	response = getLogsIndexer(fromBlock, toBlock, address, topics, limit)

	return response
}

func getLogsIndexer(fromBlock string, toBlock string, address string, topics []interface{}, limit int) rsps.GetLogsResponse {
	var err error
	var response rsps.GetLogsResponse
	response.Jsonrpc = "2.0"
	response.ID = 73

	var conditions map[string]interface{}

	var fromIndex int64
	var toIndex int64
	// Check the sequence of fromBlock and toBlock
	if fromBlock == "latest" {
		fromIndex, err = RetrieveCurrentBlockNumber()
		if err != nil {
			errMessage := err.Error()
			errors := common.Error{
				ErrorType:        1,
				ErrorDescription: errMessage,
			}
			logger.Console().Panic(errors)
			logger.File().Error(errors)
		}
	} else {
		// Convert the heximal format of fromblock to integer 64 bits type.
		from, err := historyUtilsCommon.ParseHexToInt64(fromBlock)
		if err != nil {
			errMessage := err.Error()
			errors := common.Error{
				ErrorType:        0,
				ErrorDescription: errMessage,
			}
			logger.Console().Panic(errors)
			logger.File().Error(errors)
		}
		fromIndex = from
	}

	if toBlock == "latest" {
		toIndex, err = RetrieveCurrentBlockNumber()
		if err != nil {
			errMessage := err.Error()
			errors := common.Error{
				ErrorType:        1,
				ErrorDescription: errMessage,
			}
			logger.Console().Panic(errors)
			logger.File().Error(errors)
		}
	} else {
		// Convert the heximal format of fromblock to integer 64 bits type.
		to, err := historyUtilsCommon.ParseHexToInt64(toBlock)
		if err != nil {
			errMessage := err.Error()
			errors := common.Error{
				ErrorType:        0,
				ErrorDescription: errMessage,
			}
			logger.Console().Panic(errors)
			logger.File().Error(errors)
		}
		toIndex = to
	}

	if !historyUtilsCommon.CheckLeftIsLowerThanTheRight(fromIndex, toIndex) {
		message := fmt.Sprintf("The index of fromBlock must lower than toBlock")
		errors := common.Error{
			ErrorType:        0,
			ErrorDescription: message,
		}
		logger.Console().Panic(errors)
	}

	if address == "0x0" {
		conditions = bson.M{
			"blockNumber": &bson.M{
				"$gte": fromIndex,
				"$lte": toIndex,
			},
			"logs.topics": &bson.M{
				"$all": topics,
			}}
	} else {
		conditions = bson.M{
			"blockNumber": &bson.M{
				"$gte": fromIndex,
				"$lte": toIndex,
			},
			"logs.address": address,
			"logs.topics": &bson.M{
				"$all": topics,
			},
		}
	}

	var result []receiptStrcut.Receipt

	result, err = RetrieveAllReceipt(conditions)
	if err != nil {
		errors := common.Error{
			ErrorType:        1,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
		logger.File().Error(err)
	}

	var getLogsResponses []receiptStrcut.Log

	// Set a counter to caculate how many topics we have retrieve in order to fit the limit number for request.
	var counter = 0
	for _, data := range result {

		for j := range data.Logs {
			// Notice: Because there are many topics in a block
			// Thus, when dealing with the parameter: limit, we should iterface each topics in a block.

			var getLogsResponse receiptStrcut.Log
			getLogsResponse.Address = data.Logs[j].Address
			getLogsResponse.BlockHash = data.Logs[j].BlockHash
			getLogsResponse.BlockNumber = data.Logs[j].BlockNumber
			getLogsResponse.Data = data.Logs[j].Data
			getLogsResponse.LogIndex = data.Logs[j].LogIndex
			getLogsResponse.Topics = data.Logs[j].Topics
			getLogsResponse.TransactionHash = data.Logs[j].TransactionHash
			getLogsResponse.TransactionIndex = data.Logs[j].TransactionIndex
			getLogsResponse.TransactionLogIndex = data.Logs[j].TransactionLogIndex
			getLogsResponse.Type = data.Logs[j].Type

			getLogsResponses = append(getLogsResponses, getLogsResponse)

			if limit > 0 {
				counter++
				if counter == limit {
					response.Result = getLogsResponses
					return response
				}
			}
		}
	}

	response.Result = getLogsResponses
	return response

}

// RetrieveAllReceipt retrieve all receipt data from mongo
func RetrieveAllReceipt(conditions map[string]interface{}) ([]receiptStrcut.Receipt, error) {
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
	result := []receiptStrcut.Receipt{}
	err = collection.Find(conditions).All(&result)

	if err != nil {
		message := fmt.Sprintf("Retrive receipts failded")
		err = errors.New(message)
	}

	return result, err
}

// RetrieveCurrentBlockNumber retrieve the highest blockNumber
func RetrieveCurrentBlockNumber() (int64, error) {
	var blockNumber int64
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
	var receipt []receiptStrcut.Receipt
	var maxBlockNumber = "-blockNumber"

	err = collection.Find(nil).Sort(maxBlockNumber).Limit(1).All(&receipt)

	if err != nil {
		errors := common.Error{
			ErrorType:        1,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
		logger.File().Error(err)
	}

	for _, data := range receipt {
		blockNumber = data.BlockNumber.(int64)
	}

	return blockNumber, err
}
