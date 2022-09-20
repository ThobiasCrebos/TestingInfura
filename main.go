package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	SimpleStorage "propchain_go/gen"
)

func main() {
	b, err := ioutil.ReadFile("./wallet/UTC--2022-09-20T09-14-07.385754000Z--fa39e74501a963f8a472e1160d18052d6b255c87")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial("https://polygon-mumbai.g.alchemy.com/v2/kWntfQh8xbGbzLkQVSQa10ibIhpEQ3tL")
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	chainid, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	gasprice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	add := common.HexToAddress("0x858c3647bE6CA85B308B93B6ef8d2d8ea67E9297")
	st, err := SimpleStorage.NewSimpleStorage(add, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainid)
	if err != nil {
		log.Fatal(err)
	}
	tx.GasLimit = 3000000
	tx.GasPrice = gasprice

	// transaction, err := st.Store(tx, big.NewInt(100))
	// if err != nil {
	// 	log.Fatal("err4 ", err)
	// }

	// fmt.Print(transaction.Hash())

	transaction2, err := st.Retrieve(&bind.CallOpts{})
	if err != nil {
		log.Fatal("err4 ", err)
	}

	fmt.Print("\n retrieve \n", transaction2)

	transaction3, err := st.AddPerson(tx, "dev", big.NewInt(1))
	if err != nil {
		log.Fatal("err5 ", err)
	}

	fmt.Print(" \n write \n", transaction3.Hash())

}
