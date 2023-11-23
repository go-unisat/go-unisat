package swap

import (
	"context"
	"github.com/go-unisat/go-unisat/common"
)

func GetGlobalConfig(ctx context.Context, server, bear string) (ResponseGlobalConfig, error) {
	var resp ResponseGlobalConfig
	url := GlobalConfig(server)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetBalance(ctx context.Context, server, bear, address, tick string) (ResponseBalance, error) {
	var resp ResponseBalance
	url := Balance(server, address, tick)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetDepositInfo(ctx context.Context, server, bear, address, tick string) (ResponseDepositInfo, error) {
	var resp ResponseDepositInfo
	url := DepositInfo(server, address, tick)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GetSelect(ctx context.Context, server, bear, address, query string) (ResponseSearch, error) {
	var resp ResponseSearch
	url := Select(server, address, query)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func PreDeployPool(ctx context.Context, server, bear, address, tick0, tick1 string, timestamp int64) (ResponsePreDeployPool, error) {
	var resp ResponsePreDeployPool
	url := urlPreDeployPool(server, address, tick0, tick1, timestamp)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func PreAddLiq(ctx context.Context, server, bear, address, tick0, tick1, amount0, amount1, lp, slippage string, timestamp int64) (ResponsePreAddLiq, error) {
	var resp ResponsePreAddLiq
	url := urlPreAddLiq(server, address, tick0, tick1, amount0, amount1, lp, slippage, timestamp)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func PreRemoveLiq(ctx context.Context, server, bear, address, tick0, tick1, amount0, amount1, lp, slippage string, timestamp int64) (ResponsePreRemoveLiq, error) {
	var resp ResponsePreRemoveLiq
	url := urlPreRemoveLiq(server, address, tick0, tick1, amount0, amount1, lp, slippage, timestamp)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func PreSwap(ctx context.Context, server, bear, address, tick0, tick1, amount0, amount1, slippage, exactType string, timestamp int64) (ResponsePreSwap, error) {
	var resp ResponsePreSwap
	url := urlPreSwap(server, address, tick0, tick1, amount0, amount1, slippage, exactType, timestamp)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func PoolList(ctx context.Context, server, bear, search string, offset, limit int64) (ResponsePoolList, error) {
	var resp ResponsePoolList
	url := urlPoolList(server, search, offset, limit)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func MyPoolList(ctx context.Context, server, bear, address, tick string, offset, limit int64) (ResponseMyPoolList, error) {
	var resp ResponseMyPoolList
	url := urlMyPoolList(server, address, tick, offset, limit)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func MyPool(ctx context.Context, server, bear, address, tick0, tick1 string) (ResponseMyPool, error) {
	var resp ResponseMyPool
	url := urlMyPool(server, address, tick0, tick1)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func Overview(ctx context.Context, server, bear string) (ResponseOverview, error) {
	var resp ResponseOverview
	url := urlOverview(server)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}

func GasHistory(ctx context.Context, server, bear, address string, offset, limit int64) (ResponseGasHistory, error) {
	var resp ResponseGasHistory
	url := urlGasHistory(server, address, offset, limit)
	err := common.GetWithBear(ctx, url, bear, &resp)
	return resp, err
}
