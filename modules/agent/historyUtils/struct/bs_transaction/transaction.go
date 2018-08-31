package transaction

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
