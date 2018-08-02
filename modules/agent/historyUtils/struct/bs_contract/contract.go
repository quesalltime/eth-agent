package contract

// Receipt define the struct for transaction receipt in blocks.
type Contract struct {
	BlockNumber       interface{} `bson:"blockNumber"`
	ContractAddress   string      `bson:"contractAddress"`
	CumulativeGasUsed string      `bson:"cumulativeGasUsed"`
	GasUsed           string      `bson:"gasUsed"`
	From              string      `bson:"from"`
	TransactionHash   string      `bson:"transactionHash"`
	TransactionIndex  string      `bson:"transactionIndex"`
	Timestamp         int         `bson:"timestamp"`
	Code              string      `bson:"code"`
}
