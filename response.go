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
