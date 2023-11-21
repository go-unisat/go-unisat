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

func AddressBalance(server, address string) string {
	return fmt.Sprintf("%s/v1/indexer/address/%s/balance", server, address)
}

func AddressHistory(server, address string, offset, limit int64) string {
	return fmt.Sprintf("%s/v1/indexer/address/%s/history?cursor=%d&size=%d", server, address, offset, limit)
}

func BtcUtxo(server, address string, offset, limit int64) string {
	return fmt.Sprintf("%s/v1/indexer/address/%s/utxo-data?cursor=%d&size=%d", server, address, offset, limit)
}

func InscriptionUtxo(server, address string, offset, limit int64) string {
	return fmt.Sprintf("%s/v1/indexer/address/%s/inscription-utxo-data?cursor=%d&size=%d", server, address, offset, limit)
}
