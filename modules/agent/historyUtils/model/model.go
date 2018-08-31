package model

import (
	"errors"
	"eth-agent/common"
	"eth-agent/config"

	collectionName "eth-agent/modules/agent/historyUtils/common"
	historyUtilsMongo "eth-agent/modules/agent/historyUtils/mongo"
	blockStrcut "eth-agent/modules/agent/historyUtils/struct/bs_block"
	contractStruct "eth-agent/modules/agent/historyUtils/struct/bs_contract"
	receiptStrcut "eth-agent/modules/agent/historyUtils/struct/bs_receipt"
	transactionStruct "eth-agent/modules/agent/historyUtils/struct/bs_transaction"
	"eth-agent/modules/logger"
	"fmt"
	"time"
)

var (
	dbName = config.SysConf.Mongo.DBName
)

// RetrieveBlock retrieve specific block data from mongo
// Remember we need to use the BlockWithOnlyTxHashes to mapping the block in mongodb
// Because the transaction formant in block only contain hashes...
func RetrieveBlock(conditions map[string]interface{}) ([]blockStrcut.BlockWithOnlyTxHashesIntNum, error) {
	var err error

	mongo, err := historyUtilsMongo.GetMongoSession()
	mongo.SetSocketTimeout(1 * time.Hour)
	//session.SetSocketTimeout(1 * time.Hour)

	if err != nil {
		errors := common.Error{
			ErrorType:        1,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
		logger.File().Error(err)
	}

	defer mongo.Close()

	collection := mongo.DB(dbName).C(collectionName.BsBlocks)
	result := []blockStrcut.BlockWithOnlyTxHashesIntNum{}
	err = collection.Find(conditions).All(&result)

	if err != nil {
		message := fmt.Sprintf("Retrive blocks failded, Because: %s", err)
		err = errors.New(message)
	}

	return result, err
}

// RetrieveTransactions retrieve the information of a transaction
func RetrieveTransactions(conditions map[string]interface{}) ([]transactionStruct.Transaction, error) {
	var err error

	mongo, err := historyUtilsMongo.GetMongoSession()
	mongo.SetSocketTimeout(1 * time.Hour)

	if err != nil {
		errors := common.Error{
			ErrorType:        1,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
		logger.File().Error(err)
	}

	defer mongo.Close()

	collection := mongo.DB(dbName).C(collectionName.BsTransactions)
	result := []transactionStruct.Transaction{}
	err = collection.Find(conditions).All(&result)

	if err != nil {
		fmt.Println(err)
		message := fmt.Sprintf("Retrive transaction failded")
		err = errors.New(message)
	}

	return result, err
}

// RetrieveReceipts retrieve all receipt data from mongo
func RetrieveReceipts(conditions map[string]interface{}) ([]receiptStrcut.Receipt, error) {
	var err error

	mongo, err := historyUtilsMongo.GetMongoSession()
	mongo.SetSocketTimeout(1 * time.Hour)

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

// RetrieveContracts retrieve the information of a contract
func RetrieveContracts(conditions map[string]interface{}) ([]contractStruct.Contract, error) {
	var err error

	mongo, err := historyUtilsMongo.GetMongoSession()
	mongo.SetSocketTimeout(1 * time.Hour)

	if err != nil {
		errors := common.Error{
			ErrorType:        1,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
		logger.File().Error(err)
	}

	defer mongo.Close()

	collection := mongo.DB(dbName).C(collectionName.BsContracts)
	result := []contractStruct.Contract{}
	err = collection.Find(conditions).All(&result)

	if err != nil {
		message := fmt.Sprintf("Retrive Receipt of Transaction failded")
		err = errors.New(message)
	}

	return result, err
}

// RetrieveCurrentBlockNumber retrieve the highest blockNumber
func RetrieveCurrentBlockNumber() (int64, error) {
	var blockNumber int64
	var err error

	mongo, err := historyUtilsMongo.GetMongoSession()
	mongo.SetSocketTimeout(1 * time.Hour)

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
