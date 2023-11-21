package unisat

import (
	"context"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestGetBlockchainInfo(t *testing.T) {
	bear := os.Getenv("BEAR")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	resp, err := GetBlockchainInfo(ctx, DefaultTest, bear)
	require.NoError(t, err, "GetBlockchainInfo failed")
	t.Log(resp)
}
