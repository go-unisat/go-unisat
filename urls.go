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

func InscriptionInfo(server, inscription string) string {
	return fmt.Sprintf("%s/v1/indexer/inscription/info/%s", server, inscription)
}

func BestBlockHeight(server string) string {
	return fmt.Sprintf("%s/v1/indexer/brc20/bestheight", server)
}

func Brc20List(server string, offset, limit int64) string {
	return fmt.Sprintf("%s/v1/indexer/brc20/list?start=%d&limit=%d", server, offset, limit)
}

func Brc20Info(server, ticker string) string {
	return fmt.Sprintf("%s/v1/indexer/brc20/%s/info", server, ticker)
}

func Brc20Holders(server, ticker string, offset, limit int64) string {
	return fmt.Sprintf("%s/v1/indexer/brc20/%s/holders?start=%d&limit=%d", server, ticker, offset, limit)
}

func Brc20History(server, ticker, eventType string, height, offset, limit int64) string {
	return fmt.Sprintf("%s/v1/indexer/brc20/%s/history?type=%s&height=%d&start=%d&limit=%d", server, ticker, eventType, height, offset, limit)
}

func Brc20TxHistory(server, ticker, tx, eventType string, offset, limit int64) string {
	return fmt.Sprintf("%s/v1/indexer/brc20/%s/tx/%s/history?type=%s&start=%d&limit=%d", server, ticker, tx, eventType, offset, limit)
}

func AddressBrc20Summary(server, address string, offset, limit int64) string {
	return fmt.Sprintf("%s/v1/indexer/address/%s/brc20/summary?start=%d&limit=%d", server, address, offset, limit)
}

func AddressBrc20TickInfo(server, address, ticker string) string {
	// https://open-api.unisat.io/v1/indexer/address/{address}/brc20/{ticker}/info
	return fmt.Sprintf("%s/v1/indexer/address/%s/brc20/%s/info", server, address, ticker)
}

func AddressBrc20History(server, address, ticker, eventType string, offset, limit int64) string {
	// https://open-api.unisat.io/v1/indexer/address/{address}/brc20/{ticker}/history
	return fmt.Sprintf("%s/v1/indexer/address/%s/brc20/%s/history?type=%s&start=%d&limit=%d", server, address, ticker, eventType, offset, limit)
}

func TransferableInscriptions(server, address, ticker string, offset, limit int64) string {
	// https://open-api.unisat.io/v1/indexer/address/{address}/brc20/{ticker}/transferable-inscriptions
	return fmt.Sprintf("%s/v1/indexer/address/%s/brc20/%s/transferable-inscriptions?start=%d&limit=%d", server, address, ticker, offset, limit)
}
