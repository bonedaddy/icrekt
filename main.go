package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const address = "0xb5A5F22694352C15B00323844aD545ABb2B11028"
const connURL = ".."

func main() {

	keyFile := os.Getenv("KEY_FILE")
	keyPass := os.Getenv("KEY_PASS")

	if keyFile == "" || keyPass == "" {
		log.Fatal("didnt provide KEY_FILE or KEY_PASS env")
	}

	keyFileData, err := ioutil.ReadFile(keyFile)
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(string(keyFileData)), keyPass)
	if err != nil {
		log.Fatal(err)
	}
	client, err := ethclient.Dial(connURL)
	if err != nil {
		log.Fatal(err)
	}
	contract, err := NewBindings(common.HexToAddress(address), client)
	if err != nil {
		log.Fatal(err)
	}

	for {
		// don't read pending state data
		transfersEnabled, err := contract.TokenTransfer(&bind.CallOpts{Pending: false})
		if err != nil {
			fmt.Println("error reading transfer state ", err)
			continue
		}
		// if transfers aren't enabled, then we skip
		if !transfersEnabled {
			fmt.Println("transfers already disabled")
			continue
		}
		// from here on out spends real ether
		tx, err := contract.DisableTokenTransfer(auth)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tx.Hash().String())
	}
}
