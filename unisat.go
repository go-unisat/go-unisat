package unisat

import (
	"context"
	"github.com/go-unisat/go-unisat/common"
)

func GetBlockchainInfo(ctx context.Context, server, bear string) (ResponseBlockchainInfo, error) {
	var resp ResponseBlockchainInfo
	url := BlockchainInfo(server)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetBlockTransactions(ctx context.Context, server, bear string, height, offset, limit int64) (ResponseBlockTransactions, error) {
	var resp ResponseBlockTransactions
	url := BlockTransactions(server, height, offset, limit)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetTxInfo(ctx context.Context, server, bear, tx string) (ResponseTxInfo, error) {
	var resp ResponseTxInfo
	url := TxInfo(server, tx)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetTxInputs(ctx context.Context, server, bear, tx string, offset, limit int64) (ResponseTxInputs, error) {
	var resp ResponseTxInputs
	url := TxInputs(server, tx, offset, limit)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetTxOutputs(ctx context.Context, server, bear, tx string, offset, limit int64) (ResponseTxOutputs, error) {
	var resp ResponseTxOutputs
	url := TxOutputs(server, tx, offset, limit)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetAddressBalance(ctx context.Context, server, bear, address string) (ResponseAddressBalance, error) {
	var resp ResponseAddressBalance
	url := AddressBalance(server, address)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetAddressHistory(ctx context.Context, server, bear, address string) (ResponseAddressBalance, error) {
	//var resp ResponseAddressBalance
	//url := AddressBalance(server, address)
	//err := common.GetWithBear(ctx, url, bear, &resp)
	//return resp, err
	panic("not supported")
}

func GetBtcUtxo(ctx context.Context, server, bear, address string, offset, limit int64) (ResponseBtcUtxo, error) {
	var resp ResponseBtcUtxo
	url := BtcUtxo(server, address, offset, limit)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetInscriptionUtxo(ctx context.Context, server, bear, address string, offset, limit int64) (ResponseInscriptionUtxo, error) {
	var resp ResponseInscriptionUtxo
	url := InscriptionUtxo(server, address, offset, limit)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetInscriptionInfo(ctx context.Context, server, bear, inscription string) (ResponseInscriptionInfo, error) {
	var resp ResponseInscriptionInfo
	url := InscriptionInfo(server, inscription)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}
