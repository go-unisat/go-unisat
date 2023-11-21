package main

import (
	"context"
	"encoding/json"
	"github.com/go-unisat/go-unisat"
	"log"
	"os"
	"time"
)

func main() {
	bear := os.Getenv("BEAR")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	resp, err := unisat.GetTxOutputs(ctx, unisat.DefaultMainnet, bear, "946351870f1b2005b141e58d1b548ddf11b2225c2e14cdb418f148f86b523ee1", 0, 10)
	if err != nil {
		log.Fatalf("GetTxInfo error: %s", err)
	}
	data, _ := json.MarshalIndent(resp, "", "\t")
	log.Printf("%s", string(data))
}
