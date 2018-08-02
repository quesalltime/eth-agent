package rsps

import (
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
type ReceiptReponse struct {
	Jsonrpc string          `json:"jsonrpc"`
	ID      int             `json:"id"`
	Result  receipt.Receipt `json:"result`
}
