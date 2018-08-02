package rqst

// rqst is the abbreviation of request
// define the requset format for json-rpc method.

// Payload : input data for JSON-rpc call
type Payload struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}
