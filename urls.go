package unisat

import "fmt"

const (
	DefaultMainnet = "https://open-api.unisat.io"
	DefaultTest    = "https://open-api-testnet.unisat.io"
)

func BlockchainInfo(server string) string {
	return fmt.Sprintf("%s/v1/indexer/blockchain/info", server)
}

func BlockTransactions(server string, height int64) string {
	return fmt.Sprintf("%s/v1/indexer/block/%d/txs", server, height)
}

func TxInfo(server, tx string) string {
	return fmt.Sprintf("%s/v1/indexer/tx/%s", server, tx)
}

func TxInputs(server, tx string) string {
	return fmt.Sprintf("%s/v1/indexer/tx/%s/ins", server, tx)
}

func TxOutputs(server, tx string) string {
	return fmt.Sprintf("%s/v1/indexer/tx/%s/outs", server, tx)
}
