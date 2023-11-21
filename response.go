package unisat

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
