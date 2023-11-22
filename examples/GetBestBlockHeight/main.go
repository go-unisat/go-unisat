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
	resp, err := unisat.GetBestBlockHeight(ctx, unisat.DefaultMainnet, bear)
	if err != nil {
		log.Fatalf("Request error: %s", err)
	}
	data, _ := json.MarshalIndent(resp, "", "\t")
	log.Printf("%s", string(data))
}
