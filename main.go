package main

import (
	"github.com/stellar/go/keypair"
	"log"
	"io/ioutil"
	"net/http"
	"github.com/stellar/go/clients/horizon"
)

func main() {
	pair, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Private Key:" + pair.Seed())
	log.Println("Public Key:" + pair.Address())

	resp, err := http.Get("https://horizon-testnet.stellar.org/friendbot?addr=" + pair.Address())
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

	account, err := horizon.DefaultTestNetClient.LoadAccount(pair.Address())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Balances for account:", pair.Address())

	for _, balance := range account.Balances {
		log.Println(balance)
	}
}