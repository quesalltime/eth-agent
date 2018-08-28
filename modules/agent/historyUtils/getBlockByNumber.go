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

func GetBlockByNumber(data rqst.Payload) interface{} {
	var err error
	var message string

	var blockNumberHex string
	var isNeedAllTx bool

	params := data.Params
	if len(params) != 2 {
		message = fmt.Sprintf("Invalid params: invalid length 1, expected a tuple of size 2.")
		err = errors.New(message)
		errors := common.Error{
			ErrorType:        0,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
	}
	param1 := params[0]
	param2 := params[1]

	switch value := param2.(type) {
	case bool:
		isNeedAllTx = value
	default:
		message = fmt.Sprintf("Invalid type in your params, expected a boolean.")
		err = errors.New(message)
		errors := common.Error{
			ErrorType:        0,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
	}

	blockNumberHex = param1.(string)
	err = historyUtilsCommon.CheckBlockNumberFormat(blockNumberHex)
	if err != nil {
		errors := common.Error{
			ErrorType:        0,
			ErrorDescription: err.Error(),
		}
		logger.Console().Panic(errors)
	}

	blockNumber, err := historyUtilsCommon.ParseHexToInt64(blockNumberHex)

	fmt.Println("================getBlockByNumber initial parameter================")
	fmt.Printf("blockNumber:%s, isNeedAllTx: %t", blockNumber, isNeedAllTx)

	response := getBlockByNumberIndexer(blockNumber, isNeedAllTx)
	var responseBlockType rsps.GetBlockWithOnlyTxHashesResponse
	var responseBlockTxHashOnlyType rsps.GetBlockOnlyTxHashResponse
	var responseEmpty rsps.EmptyResponse

	switch target := response.(type) {
	case rsps.GetBlockOnlyTxHashResponse:
		responseBlockTxHashOnlyType = response.(rsps.GetBlockOnlyTxHashResponse)
		return responseBlockTxHashOnlyType
	case rsps.GetBlockWithOnlyTxHashesResponse:
		responseBlockType = response.(rsps.GetBlockWithOnlyTxHashesResponse)
		return responseBlockType
	case rsps.EmptyResponse:
		responseEmpty = response.(rsps.EmptyResponse)
		return responseEmpty
	default:
		logger.Console().Panic(fmt.Sprintf("The type of the parameter: %v ", target))
	}

	return 0
}

func getBlockByNumberIndexer(blockNumber int64, isNeedAllTx bool) interface{} {
	var response rsps.EmptyResponse
	var responseBlockType rsps.GetBlockWithOnlyTxHashesResponse
	var responseBlockTxHashOnlyType rsps.GetBlockOnlyTxHashResponse

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

	if len(result) == 0 {
		response.Jsonrpc = "2.0"
		response.ID = 73
		response.Result = ""
		return response
	}

	if isNeedAllTx {
		responseBlockType.Jsonrpc = "2.0"
		responseBlockType.ID = 73
		responseBlockType.Jsonrpc = "2.0"
		responseBlockType.ID = 73
		responseBlockType.Result.Author = result[0].Author
		responseBlockType.Result.Difficulty = result[0].Difficulty
		responseBlockType.Result.ExtraData = result[0].ExtraData
		responseBlockType.Result.GasLimit = result[0].GasLimit
		responseBlockType.Result.GasUsed = result[0].GasUsed
		responseBlockType.Result.LogsBloom = result[0].LogsBloom
		responseBlockType.Result.Miner = result[0].Miner
		responseBlockType.Result.MixHash = result[0].MixHash
		responseBlockType.Result.Nonce = result[0].Nonce
		responseBlockType.Result.ParentHash = result[0].ParentHash
		responseBlockType.Result.ReceiptsRoot = result[0].ReceiptsRoot
		responseBlockType.Result.SealFields = result[0].SealFields
		responseBlockType.Result.Sha3Uncles = result[0].Sha3Uncles
		responseBlockType.Result.Size = result[0].Size
		responseBlockType.Result.StateRoot = result[0].StateRoot
		responseBlockType.Result.Timestamp = result[0].Timestamp
		responseBlockType.Result.TotalDifficulty = result[0].TotalDifficulty
		responseBlockType.Result.Hash = result[0].Hash
		responseBlockType.Result.TransactionsRoot = result[0].TransactionsRoot
		responseBlockType.Result.Uncles = result[0].Uncles
		responseBlockType.Result.Transactions = result[0].Transactions

		responseBlockType.Result.Number = historyUtilsCommon.ParseInt64ToHex(result[0].Number)
		return responseBlockType
	}

	responseBlockTxHashOnlyType.Jsonrpc = "2.0"
	responseBlockTxHashOnlyType.ID = 73
	responseBlockTxHashOnlyType.Result.Author = result[0].Author
	responseBlockTxHashOnlyType.Result.Difficulty = result[0].Difficulty
	responseBlockTxHashOnlyType.Result.ExtraData = result[0].ExtraData
	responseBlockTxHashOnlyType.Result.GasLimit = result[0].GasLimit
	responseBlockTxHashOnlyType.Result.GasUsed = result[0].GasUsed
	responseBlockTxHashOnlyType.Result.LogsBloom = result[0].LogsBloom
	responseBlockTxHashOnlyType.Result.Miner = result[0].Miner
	responseBlockTxHashOnlyType.Result.MixHash = result[0].MixHash
	responseBlockTxHashOnlyType.Result.Nonce = result[0].Nonce
	responseBlockTxHashOnlyType.Result.ParentHash = result[0].ParentHash
	responseBlockTxHashOnlyType.Result.ReceiptsRoot = result[0].ReceiptsRoot
	responseBlockTxHashOnlyType.Result.SealFields = result[0].SealFields
	responseBlockTxHashOnlyType.Result.Sha3Uncles = result[0].Sha3Uncles
	responseBlockTxHashOnlyType.Result.Size = result[0].Size
	responseBlockTxHashOnlyType.Result.StateRoot = result[0].StateRoot
	responseBlockTxHashOnlyType.Result.Timestamp = result[0].Timestamp
	responseBlockTxHashOnlyType.Result.TotalDifficulty = result[0].TotalDifficulty
	responseBlockTxHashOnlyType.Result.Hash = result[0].Hash
	responseBlockTxHashOnlyType.Result.TransactionsRoot = result[0].TransactionsRoot
	responseBlockTxHashOnlyType.Result.Uncles = result[0].Uncles

	// block number
	blockNumberHex := historyUtilsCommon.ParseInt64ToHex(result[0].Number)
	responseBlockTxHashOnlyType.Result.Number = blockNumberHex

	/*
		// tx hashes
		txs := result[0].Transactions
		txsHash := make([]string, len(txs))
		for i, tx := range txs {
			txsHash[i] = tx.Hash
		}
		responseBlockTxHashOnlyType.Result.Transactions = txsHash
	*/
	responseBlockTxHashOnlyType.Result.Transactions = result[0].Transactions
	return responseBlockTxHashOnlyType

}
