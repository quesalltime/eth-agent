package block

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
	Number           int    `bson:"number"`
	ParentHash       string `bson:"parentHash"`
	ReceiptsRoot     string `bson:"receiptsRoot"`
	SealFields       string `bson:"sealFields"`
	Sha3Uncles       string `bson:"sha3Uncles"`
	Size             string `bson:"size"`
	StateRoot        string `bson:"stateRoot"`
	Timestamp        int    `bson:"timestamp"`
	TotalDifficulty  string `bson:"totalDifficulty"`
	Transactions     []Transaction
	TransactionsRoot string `bson:"transactionsRoot"`
	Uncles           string `bson:"uncles"`
}

type Transaction struct {
	BlockHash        string `bson:"blockHash"`
	BlockNumber      string `bson:"blockNumber"`
	ChainID          string `bson:"chainId"`
	Condition        string `bson:"condition"`
	Creates          string `bson:"creates"`
	From             string `bson:"from"`
	Gas              string `bson:"gas"`
	GasPrice         string `bson:"gasPrice"`
	Hash             string `bson:"hash"`
	Input            string `bson:"input"`
	Nonce            string `bson:"nonce"`
	PublicKey        string `bson:"publicKey"`
	Raw              string `bson:"raw"`
	S                string `bson:"s"`
	R                string `bson:"r"`
	V                string `bson:"v"`
	StandardV        string `bson:"standardV"`
	To               string `bson:"to"`
	TransactionIndex string `bson:"transactionIndex"`
	Value            string `bson:"value"`
	Timestamp        int    `bson:"timestamp"`
}
