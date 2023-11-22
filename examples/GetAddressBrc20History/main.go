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
	resp, err := unisat.GetAddressBrc20History(
		ctx, unisat.DefaultMainnet, bear,
		"bc1pxgpyskz606rlpfjsg3jsluuqedxflyrnfdscftcc7slesqlgqxfsa2mvcf",
		"3518", "mint",
		0, 10,
	)
	if err != nil {
		log.Fatalf("Request error: %s", err)
	}
	data, _ := json.MarshalIndent(resp, "", "  ")
	log.Printf("%s", string(data))
}
