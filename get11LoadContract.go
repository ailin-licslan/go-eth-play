package main

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	store "go-ethereum/store"
)

const (

	//合约地址       0xf35bFE94C64d1dF4C86F0488ff687e307a2Ce3c6
	contractAddr = "0x8D4141ec2b522dE5Cf42705C3010541B4B3EC24e"
)

// get11LoadContract
func get11LoadContract() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/00fdc4fc92d945ea9e8b5a0157764xxx")
	if err != nil {
		log.Fatal(err)
	}
	storeContract, err := store.NewStore(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}

	_ = storeContract
	println(storeContract)
}
