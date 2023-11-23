package swap

import "fmt"

func GlobalConfig(server string) string {
	// https://open-api.unisat.io/v1/brc20-swap/config
	return fmt.Sprintf("%s/v1/brc20-swap/config", server)
}

func Balance(server, address, tick string) string {
	// https://open-api.unisat.io/v1/brc20-swap/balance
	return fmt.Sprintf("%s/v1/brc20-swap/balance?address=%s&tick=%s", server, address, tick)
}

func DepositInfo(server, address, tick string) string {
	// https://open-api.unisat.io/v1/brc20-swap/deposit_info
	return fmt.Sprintf("%s/v1/brc20-swap/deposit_info?address=%s&tick=%s", server, address, tick)
}

func Select(server, address, query string) string {
	// https://open-api.unisat.io/v1/brc20-swap/select
	return fmt.Sprintf("%s/v1/brc20-swap/select?address=%s&search=%s", server, address, query)
}

func urlPreDeployPool(server, address, tick0, tick1 string, timestamp int64) string {
	// https://open-api.unisat.io/v1/brc20-swap/pre_deploy_pool
	return fmt.Sprintf("%s/v1/brc20-swap/pre_deploy_pool?address=%s&tick0=%s&tick1=%s&ts=%d", server, address, tick0, tick1, timestamp)
}

func urlPreAddLiq(server, address, tick0, tick1, amount0, amount1, lp, slippage string, timestamp int64) string {
	// https://open-api.unisat.io/v1/brc20-swap/pre_add_liq
	return fmt.Sprintf(
		"%s/v1/brc20-swap/pre_add_liq?address=%s&tick0=%s&tick1=%s&amount0=%s&amount1=%s&lp=%s&slippage=%s&ts=%d",
		server, address, tick0, tick1, amount0, amount1, lp, slippage, timestamp,
	)
}

func urlPreRemoveLiq(server, address, tick0, tick1, amount0, amount1, lp, slippage string, timestamp int64) string {
	// https://open-api.unisat.io/v1/brc20-swap/pre_remove_liq
	return fmt.Sprintf(
		"%s/v1/brc20-swap/pre_remove_liq?address=%s&tick0=%s&tick1=%s&amount0=%s&amount1=%s&lp=%s&slippage=%s&ts=%d",
		server, address, tick0, tick1, amount0, amount1, lp, slippage, timestamp,
	)
}

func urlPreSwap(server, address, tickIn, tickOut, amountIn, amountOut, slippage, exactType string, timestamp int64) string {
	// https://open-api.unisat.io/v1/brc20-swap/pre_swap
	return fmt.Sprintf(
		"%s/v1/brc20-swap/pre_swap?address=%s&tickIn=%s&tickOut=%s&amountIn=%s&amountOut=%s&slippage=%s&exactType=%s&ts=%d",
		server, address, tickIn, tickOut, amountIn, amountOut, slippage, exactType, timestamp,
	)
}

func urlPoolList(server, search string, offset, limit int64) string {
	// https://open-api.unisat.io/v1/brc20-swap/pool_list
	return fmt.Sprintf("%s/v1/brc20-swap/pool_list?search=%s&start=%d&limit=%d", server, search, offset, limit)
}

func urlMyPoolList(server, address, tick string, offset, limit int64) string {
	// https://open-api.unisat.io/v1/brc20-swap/my_pool_list
	return fmt.Sprintf("%s/v1/brc20-swap/my_pool_list?address=%s&tick=%s&start=%d&limit=%d", server, address, tick, offset, limit)
}

func urlMyPool(server, address, tick0, tick1 string) string {
	// https://open-api.unisat.io/v1/brc20-swap/my_pool
	return fmt.Sprintf("%s/v1/brc20-swap/my_pool?address=%s&tick0=%s&tick1=%s", server, address, tick0, tick1)
}

func urlOverview(server string) string {
	// https://open-api.unisat.io/v1/brc20-swap/overview
	return fmt.Sprintf("%s/v1/brc20-swap/overview", server)
}

func urlGasHistory(server, address string, offset, limit int64) string {
	// https://open-api.unisat.io/v1/brc20-swap/gas_history
	return fmt.Sprintf("%s/v1/brc20-swap/gas_history?address=%s&start=%d&limit=%d", server, address, offset, limit)
}
