package unisat

import "fmt"

const (
	DefaultMainnet = "https://open-api.unisat.io"
	DefaultTest    = "https://open-api-testnet.unisat.io"
)

func BlockchainInfo(server string) string {
	return fmt.Sprintf("%s/v1/indexer/blockchain/info", server)
}

func BlockTransactions(server string, height, offset, limit int64) string {
	return fmt.Sprintf("%s/v1/indexer/block/%d/txs?cursor=%d&size=%d", server, height, offset, limit)
}

func TxInfo(server, tx string) string {
	return fmt.Sprintf("%s/v1/indexer/tx/%s", server, tx)
}

func TxInputs(server, tx string, offset, limit int64) string {
	return fmt.Sprintf("%s/v1/indexer/tx/%s/ins?cursor=%d&size=%d", server, tx, offset, limit)
}

func TxOutputs(server, tx string, offset, limit int64) string {
	return fmt.Sprintf("%s/v1/indexer/tx/%s/outs?cursor=%d&size=%d", server, tx, offset, limit)
}
