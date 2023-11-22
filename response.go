package unisat

import "math/big"

type DataBlockchainInfo struct {
	Chain         string `json:"chain"`
	Blocks        int64  `json:"blocks"`
	Headers       int64  `json:"headers"`
	BestBlockHash string `json:"bestBlockHash"`
	PrevBlockHash string `json:"prevBlockHash"`
	Difficulty    string `json:"difficulty"`
	MedianTime    int64  `json:"medianTime"`
	ChainWork     string `json:"chainwork"`
}

type ResponseBlockchainInfo struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data DataBlockchainInfo `json:"data"`
}

type Tx struct {
	TxId             string   `json:"txid"`
	Ins              int      `json:"nIn"`
	Outs             int      `json:"nOut"`
	Size             int      `json:"size"`
	WitOffset        int      `json:"witOffset"`
	Locktime         int      `json:"locktime"`
	InSatoshi        *big.Int `json:"inSatoshi"`
	OutSatoshi       *big.Int `json:"outSatoshi"`
	NewInscriptions  int      `json:"nNewInscription"`
	InInscriptions   int      `json:"nInInscription"`
	OutInscriptions  int      `json:"nOutInscription"`
	LostInscriptions int      `json:"nLostInscription"`
	Timestamp        int64    `json:"timestamp"`
	Height           int64    `json:"height"`
	BlockId          string   `json:"blkid"`
	Index            int      `json:"idx"`
	Confirmations    int      `json:"confirmations"`
}

type ResponseBlockTransactions struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data []Tx `json:"data"`
}

type ResponseTxInfo struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data Tx `json:"data"`
}

type Inscription struct {
	InscriptionNumber int64  `json:"inscriptionNumber"`
	InscriptionId     string `json:"inscriptionId"`
	Offset            int    `json:"offset"`
	Moved             bool   `json:"moved"`
	IsBRC20           bool   `json:"isBRC20"`
}

type Input struct {
	Height       int64         `json:"height"`
	TxId         string        `json:"txid"`
	Index        int           `json:"idx"`
	ScriptSig    string        `json:"scriptSig"`
	ScriptWits   string        `json:"scriptWits"`
	Sequence     int           `json:"sequence"`
	HeightTxo    int64         `json:"heightTxo"`
	Utxid        string        `json:"utxid"`
	Vout         int           `json:"vout"`
	Address      string        `json:"address"`
	CodeType     int           `json:"codeType"`
	Satoshi      *big.Int      `json:"satoshi"`
	ScriptType   string        `json:"scriptType"`
	ScriptPk     string        `json:"scriptPk"`
	Inscriptions []Inscription `json:"inscriptions"`
}

type ResponseTxInputs struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data []Input `json:"data"`
}

type Output struct {
	TxId         string        `json:"txid"`
	Vout         int           `json:"vout"`
	Address      string        `json:"address"`
	CodeType     int           `json:"codeType"`
	Satoshi      *big.Int      `json:"satoshi"`
	ScriptType   string        `json:"scriptType"`
	ScriptPk     string        `json:"scriptPk"`
	Height       int64         `json:"height"`
	Index        int           `json:"idx"`
	Inscriptions []Inscription `json:"inscriptions"`
	TxSpent      string        `json:"txidSpent"`
	HeightSpent  int64         `json:"heightSpent"`
}

type ResponseTxOutputs struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data []Output `json:"data"`
}

type Balance struct {
	Address string `json:"address"`

	Satoshi        *big.Int `json:"satoshi"`
	PendingSatoshi *big.Int `json:"pendingSatoshi"`
	UtxoCount      int64    `json:"utxoCount"`

	BtcSatoshi        *big.Int `json:"btcSatoshi"`
	BtcPendingSatoshi *big.Int `json:"btcPendingSatoshi"`
	BtcUtxoCount      int64    `json:"btcUtxoCount"`

	InscriptionSatoshi        *big.Int `json:"inscriptionSatoshi"`
	InscriptionPendingSatoshi *big.Int `json:"inscriptionPendingSatoshi"`
	InscriptionUtxoCount      int64    `json:"inscriptionUtxoCount"`
}

type ResponseAddressBalance struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data Balance `json:"data"`
}

type UTXO struct {
	TxId         string        `json:"txid"`
	Vout         int           `json:"vout"`
	Satoshi      *big.Int      `json:"satoshi"`
	ScriptType   string        `json:"scriptType"`
	ScriptPk     string        `json:"scriptPk"`
	CodeType     int           `json:"codeType"`
	Address      string        `json:"address"`
	Height       int64         `json:"height"`
	Index        int           `json:"idx"`
	IsOpInRBF    bool          `json:"isOpInRBF"`
	Inscriptions []Inscription `json:"inscriptions"`
}

type DataUtxoList struct {
	Cursor                int    `json:"cursor"`
	Total                 int    `json:"total"`
	TotalConfirmed        int    `json:"totalConfirmed"`
	TotalUnconfirmed      int    `json:"totalUnconfirmed"`
	TotalUnconfirmedSpent int    `json:"totalUnconfirmedSpent"`
	Utxo                  []UTXO `json:"utxo"`
}

type ResponseUtxoList struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data DataUtxoList `json:"data"`
}

type ResponseBtcUtxo ResponseUtxoList
type ResponseInscriptionUtxo ResponseUtxoList

type DataInscriptionInfo struct {
	Utxo              UTXO        `json:"utxo"`
	Address           string      `json:"address"`
	Offset            int         `json:"offset"`
	InscriptionIndex  int         `json:"inscriptionIndex"`
	InscriptionNumber int64       `json:"inscriptionNumber"`
	InscriptionId     string      `json:"inscriptionId"`
	ContentType       string      `json:"contentType"`
	ContentLength     int         `json:"contentLength"`
	ContentBody       string      `json:"contentBody"`
	Height            int64       `json:"height"`
	Timestamp         int64       `json:"timestamp"`
	InSatoshi         *big.Int    `json:"inSatoshi"`
	OutSatoshi        *big.Int    `json:"outSatoshi"`
	Brc20             interface{} `json:"brc20"`
	Detail            interface{} `json:"detail"`
}

type ResponseInscriptionInfo struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data DataInscriptionInfo `json:"data"`
}

type DataBestHeight struct {
	Height    int64  `json:"height"`
	BlockId   string `json:"blockid"`
	Timestamp int64  `json:"timestamp"`
	Total     int64  `json:"total"`
}

type ResponseBestHeight struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data DataBestHeight `json:"data"`
}

type DataBrc20List struct {
	Height int64    `json:"height"`
	Total  int64    `json:"total"`
	Start  int64    `json:"start"`
	Detail []string `json:"detail"`
}

type ResponseBrc20List struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data DataBrc20List `json:"data"`
}

type DataBrc20Info struct {
	Ticker        string `json:"ticker"`
	Creator       string `json:"creator"`
	InscriptionId string `json:"inscriptionId"`
	Max           string `json:"max"`
	Decimal       int    `json:"decimal"`
	Limit         string `json:"limit"`
	TxId          string `json:"txid"`

	InscriptionNumber      int64 `json:"inscriptionNumber"`
	InscriptionNumberStart int64 `json:"inscriptionNumberStart"`
	InscriptionNumberEnd   int64 `json:"inscriptionNumberEnd"`
	DeployHeight           int64 `json:"deployHeight"`
	DeployBlocktime        int64 `json:"deployBlocktime"`
	CompleteHeight         int64 `json:"completeHeight"`
	CompleteBlocktime      int64 `json:"completeBlocktime"`

	HoldersCount int64 `json:"holdersCount"`
	HistoryCount int64 `json:"historyCount"`
	MintTimes    int   `json:"mintTimes"`

	Minted             string `json:"minted"`
	TotalMinted        string `json:"totalMinted"`
	ConfirmedMinted    string `json:"confirmedMinted"`
	ConfirmedMinted1h  string `json:"confirmedMinted1h"`
	ConfirmedMinted24h string `json:"confirmedMinted24h"`
}

type ResponseBrc20Info struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data DataBrc20Info `json:"data"`
}

type Holder struct {
	Address             string `json:"address"`
	OverallBalance      string `json:"overallBalance"`
	TransferableBalance string `json:"transferableBalance"`
	AvailableBalance    string `json:"availableBalance"`
}

type DataBrc20Holders struct {
	Height int64    `json:"height"`
	Total  int64    `json:"total"`
	Start  int64    `json:"start"`
	Detail []Holder `json:"detail"`
}

type ResponseBrc20Holders struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data DataBrc20Holders `json:"data"`
}

type Event struct {
	Ticker            string   `json:"ticker"`
	Type              string   `json:"type"`
	Valid             bool     `json:"valid"`
	TxId              string   `json:"txid"`
	Index             int      `json:"idx"`
	Vout              int      `json:"vout"`
	InscriptionNumber int64    `json:"inscriptionNumber"`
	InscriptionId     string   `json:"inscriptionId"`
	From              string   `json:"from"`
	To                string   `json:"to"`
	Satoshi           *big.Int `json:"satoshi"`
	Amount            string   `json:"amount"`

	OverallBalance   string `json:"overallBalance"`
	TransferBalance  string `json:"transferBalance"`
	AvailableBalance string `json:"availableBalance"`

	Height    int64  `json:"height"`
	TxIndex   int    `json:"txidx"`
	Blockhash string `json:"blockhash"`
	Blocktime int64  `json:"blocktime"`
}

type DataBrc20History struct {
	Height int64   `json:"height"`
	Total  int64   `json:"total"`
	Start  int64   `json:"start"`
	Detail []Event `json:"detail"`
}

type ResponseBrc20History struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data DataBrc20History `json:"data"`
}

type ResponseBrc20TxHistory ResponseBrc20History

type Brc20Balance struct {
	Ticker              string `json:"ticker"`
	OverallBalance      string `json:"overallBalance"`
	TransferableBalance string `json:"transferableBalance"`
	AvailableBalance    string `json:"availableBalance"`
}

type DataBrc20Summary struct {
	Height int64 `json:"height"`
	Total  int64 `json:"total"`
	Start  int64 `json:"start"`

	Detail []Brc20Balance `json:"detail"`
}

type ResponseBrc20Summary struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data DataBrc20Summary `json:"data"`
}

type HistoryData struct {
	Op     string `json:"op"`
	Tick   string `json:"tick"`
	Amt    string `json:"amt"`
	Minted string `json:"minted"`
}

type InscriptionHistory struct {
	Data              HistoryData `json:"data"`
	InscriptionNumber int64       `json:"inscriptionNumber"`
	InscriptionId     string      `json:"inscriptionId"`
	Confirmations     int64       `json:"confirmations"`
}

type DataAddressTickInfo struct {
	Ticker string `json:"ticker"`

	OverallBalance         string `json:"overallBalance"`
	TransferableBalance    string `json:"transferableBalance"`
	AvailableBalance       string `json:"availableBalance"`
	AvailableBalanceSafe   string `json:"availableBalanceSafe"`
	AvailableBalanceUnSafe string `json:"availableBalanceUnSafe"`

	TransferableCount int `json:"transferableCount"`

	// TODO: check this structure
	TransferableInscriptions []interface{} `json:"transferableInscriptions"`

	HistoryCount int `json:"historyCount"`

	HistoryInscriptions []InscriptionHistory `json:"historyInscriptions"`
}

type ResponseAddressBrc20TickInfo struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`

	Data DataAddressTickInfo `json:"data"`
}

type ResponseAddressBrc20History ResponseBrc20History
type ResponseTransferableInscriptions ResponseBrc20History
