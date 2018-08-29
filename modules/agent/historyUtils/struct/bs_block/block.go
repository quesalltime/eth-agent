package block

import (
	transaction "eth-agent/modules/agent/historyUtils/struct/bs_transaction"
)

type Block struct {
	Author           string `bson:"author"`
	Difficulty       string `bson:"difficulty"`
	ExtraData        string `bson:"extraData"`
	GasLimit         string `bson:"gasLimit"`
	GasUsed          string `bson:"gasUsed"`
	Hash             string `bson:"hash"`
	LogsBloom        string `bson:"logsBloom"`
	Miner            string `bson:"miner"`
	MixHash          string `bson:"mixHash"`
	Nonce            string `bson:"nonce"`
	Number           string `bson:"number"`
	ParentHash       string `bson:"parentHash"`
	ReceiptsRoot     string `bson:"receiptsRoot"`
	SealFields       string `bson:"sealFields"`
	Sha3Uncles       string `bson:"sha3Uncles"`
	Size             string `bson:"size"`
	StateRoot        string `bson:"stateRoot"`
	Timestamp        int    `bson:"timestamp"`
	TotalDifficulty  string `bson:"totalDifficulty"`
	Transactions     []transaction.Transaction
	TransactionsRoot string `bson:"transactionsRoot"`
	Uncles           string `bson:"uncles"`
}

type BlockWithOnlyTxHashes struct {
	Author           string `bson:"author"`
	Difficulty       string `bson:"difficulty"`
	ExtraData        string `bson:"extraData"`
	GasLimit         string `bson:"gasLimit"`
	GasUsed          string `bson:"gasUsed"`
	Hash             string `bson:"hash"`
	LogsBloom        string `bson:"logsBloom"`
	Miner            string `bson:"miner"`
	MixHash          string `bson:"mixHash"`
	Nonce            string `bson:"nonce"`
	Number           string `bson:"number"`
	ParentHash       string `bson:"parentHash"`
	ReceiptsRoot     string `bson:"receiptsRoot"`
	SealFields       string `bson:"sealFields"`
	Sha3Uncles       string `bson:"sha3Uncles"`
	Size             string `bson:"size"`
	StateRoot        string `bson:"stateRoot"`
	Timestamp        int    `bson:"timestamp"`
	TotalDifficulty  string `bson:"totalDifficulty"`
	Transactions     []string
	TransactionsRoot string `bson:"transactionsRoot"`
	Uncles           string `bson:"uncles"`
}

type BlockWithOnlyTxHashesIntNum struct {
	Author           string `bson:"author"`
	Difficulty       string `bson:"difficulty"`
	ExtraData        string `bson:"extraData"`
	GasLimit         string `bson:"gasLimit"`
	GasUsed          string `bson:"gasUsed"`
	Hash             string `bson:"hash"`
	LogsBloom        string `bson:"logsBloom"`
	Miner            string `bson:"miner"`
	MixHash          string `bson:"mixHash"`
	Nonce            string `bson:"nonce"`
	Number           int    `bson:"number"`
	ParentHash       string `bson:"parentHash"`
	ReceiptsRoot     string `bson:"receiptsRoot"`
	SealFields       string `bson:"sealFields"`
	Sha3Uncles       string `bson:"sha3Uncles"`
	Size             string `bson:"size"`
	StateRoot        string `bson:"stateRoot"`
	Timestamp        int    `bson:"timestamp"`
	TotalDifficulty  string `bson:"totalDifficulty"`
	Transactions     []string
	TransactionsRoot string `bson:"transactionsRoot"`
	Uncles           string `bson:"uncles"`
}
