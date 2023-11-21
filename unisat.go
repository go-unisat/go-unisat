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
