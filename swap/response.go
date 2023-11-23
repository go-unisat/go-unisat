package swap

type DataConfig struct {
	ModuleId                  string `json:"moduleId"`
	ServiceGasTick            string `json:"serviceGasTick"`
	PendingDepositDirectNum   int    `json:"pendingDepositDirectNum"`
	PendingDepositMatchingNum int    `json:"pendingDepositMatchingNum"`
}

type ResponseGlobalConfig struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data DataConfig `json:"data"`
}

type SwapBalance struct {
	Module           string `json:"module"`
	Swap             string `json:"swap"`
	PendingSwap      string `json:"pendingSwap"`
	PendingAvailable string `json:"pendingAvailable"`
}

type DataBalance struct {
	Balance SwapBalance `json:"balance"`
	Decimal string      `json:"decimal"`
}

type ResponseBalance struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data DataBalance `json:"data"`
}

type DataDepositInfo struct {
	DailyAmount      string `json:"dailyAmount"`
	DailyLimit       string `json:"dailyLimit"`
	RecommendDeposit string `json:"recommendDeposit"`
}

type ResponseDepositInfo struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data DataDepositInfo `json:"data"`
}

type DataSearch struct {
	Tick         string `json:"tick"`
	Brc20Balance string `json:"brc20Balance"`
	SwapBalance  string `json:"swapBalance"`
	Decimal      string `json:"decimal"`
}

type ResponseSearch struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data []DataSearch `json:"data"`
}

type DataPreDeployPool struct {
	SignMsg            string  `json:"signMsg"`
	BytesL1            float64 `json:"bytesL1"`
	BytesL2            float64 `json:"bytesL2"`
	FeeRate            string  `json:"feeRate"`
	GasPrice           string  `json:"gasPrice"`
	ServiceFeeL1       string  `json:"serviceFeeL1"`
	ServiceFeeL2       string  `json:"serviceFeeL2"`
	UnitUsdPriceL1     string  `json:"unitUsdPriceL1"`
	UnitUsdPriceL2     string  `json:"unitUsdPriceL2"`
	ServiceTickBalance string  `json:"serviceTickBalance"`
}

type ResponsePreDeployPool struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data DataPreDeployPool `json:"data"`
}

type ResponsePreAddLiq ResponsePreDeployPool
type ResponsePreRemoveLiq ResponsePreDeployPool
type ResponsePreSwap ResponsePreDeployPool

type Pool struct {
	Tick0     string `json:"tick0"`
	Tick1     string `json:"tick1"`
	Tvl       string `json:"tvl"`
	Volume24h string `json:"volume24h"`
	Volume7d  string `json:"volume7d"`
	Lp        string `json:"lp"`
}

type DataPoolList struct {
	Total int    `json:"total"`
	List  []Pool `json:"list"`
}

type ResponsePoolList struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data DataPoolList `json:"data"`
}

type PoolInfo struct {
	Lp          string `json:"lp"`
	ShareOfPool string `json:"shareOfPool"`
	Tick0       string `json:"tick0"`
	Tick1       string `json:"tick1"`
	Amount0     string `json:"amount0"`
	Amount1     string `json:"amount1"`
}

type DataPoolInfo struct {
	Total int        `json:"total"`
	List  []PoolInfo `json:"list"`
}

type ResponseMyPoolList struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data []DataPoolInfo `json:"data"`
}

type ResponseMyPool struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data DataPoolInfo `json:"data"`
}

type DataOverview struct {
	Liquidity    string `json:"liquidity"`
	Volume7d     string `json:"volume7d"`
	Volume24h    string `json:"volume24h"`
	Transactions int    `json:"transactions"`
	Pairs        int    `json:"pairs"`
}

type ResponseOverview struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data DataOverview `json:"data"`
}

type HistoryGas struct {
	FuncType string `json:"funcType"`
	TickA    string `json:"tickA"`
	TickB    string `json:"tickB"`
	Gas      string `json:"gas"`
	Ts       int64  `json:"ts"`
}

type DataGasHistory struct {
	Total int          `json:"total"`
	List  []HistoryGas `json:"list"`
}

type ResponseGasHistory struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data DataGasHistory `json:"data"`
}
