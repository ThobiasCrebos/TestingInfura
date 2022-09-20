package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var url = "https://polygon-mumbai.g.alchemy.com/v2/kWntfQh8xbGbzLkQVSQa10ibIhpEQ3tL"

func main() {
	client, err := ethclient.DialContext(context.Background(), url)

	if err != nil {
		log.Fatalf("error %v", err)
	}

	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)

	if err != nil {
		log.Fatalf("error 2 %v", err)
	}

	fmt.Print(block.Number())

}
