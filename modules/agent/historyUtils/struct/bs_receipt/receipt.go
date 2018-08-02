package receipt

// Receipt define the struct for transaction receipt in blocks.
type Receipt struct {
	BlockHash         string      `bson:"blockHash"`
	BlockNumber       interface{} `bson:"blockNumber"`
	ContractAddress   string      `bson:"contractAddress"`
	CumulativeGasUsed string      `bson:"cumulativeGasUsed"`
	GasUsed           string      `bson:"gasUsed"`
	Logs              []Log
	LogsBloom         string `bson:"logsBloom"`
	Root              string `bson:"root"`
	Status            string `bson:"status"`
	TransactionHash   string `bson:"transactionHash"`
	TransactionIndex  string `bson:"transactionIndex"`
	Timestamp         int    `bson:"timestamp"`
}

// Log define the struct for logs in transaction receipt .
type Log struct {
	Address             string   `bson:"address"`
	BlockHash           string   `bson:"blockHash"`
	BlockNumber         string   `bson:"blockNumber"`
	Data                string   `bson:"data"`
	LogIndex            string   `bson:"logIndex"`
	Topics              []string `bson:"topics"`
	TransactionHash     string   `bson:"transactionHash"`
	TransactionIndex    string   `bson:"transactionIndex"`
	TransactionLogIndex string   `bson:"transactionLogIndex"`
	Type                string   `bson:"type"`
	// Timestamp           int      `bson:"timestamp"`
}
