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
	resp, err := unisat.GetInscriptionUtxo(ctx, unisat.DefaultMainnet, bear, "bc1pxgpyskz606rlpfjsg3jsluuqedxflyrnfdscftcc7slesqlgqxfsa2mvcf", 0, 10)
	if err != nil {
		log.Fatalf("Request error: %s", err)
	}
	data, _ := json.MarshalIndent(resp, "", "\t")
	log.Printf("%s", string(data))
}
