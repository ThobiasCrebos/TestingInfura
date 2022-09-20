package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "password"
	a, err := key.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(a.Address)

	b, err := ioutil.ReadFile("./wallet/UTC--2022-09-20T09-14-07.385754000Z--fa39e74501a963f8a472e1160d18052d6b255c87")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(b)

}
