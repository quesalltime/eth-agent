package rsps

import (
	block "eth-agent/modules/agent/historyUtils/struct/bs_block"
	receipt "eth-agent/modules/agent/historyUtils/struct/bs_receipt"
)

// GetLogsResponse return the response of the json-rpc request: getLogs
type GetLogsResponse struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      int           `json:"id"`
	Result  []receipt.Log `json:"result"`
}

// GetBlockTransactionCountByNumber response format
type GBTCResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

// Transaction(tx) response format
type ReceiptResponse struct {
	Jsonrpc string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  receipt.Receipt `json:"result`
}

// GetCode response format
type GetCodeResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

type GetBlockResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Result  block.Block `json:"result"`
}

type GetBlockWithOnlyTxHashesResponse struct {
	Jsonrpc string                      `json:"jsonrpc"`
	ID      int                         `json:"id"`
	Result  block.BlockWithOnlyTxHashes `json:"result"`
}

type EmptyResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}
