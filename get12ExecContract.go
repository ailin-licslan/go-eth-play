package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	store "go-ethereum/store"
	"log"
	"math/big"
)

const (

	//合约地址       0xf35bFE94C64d1dF4C86F0488ff687e307a2Ce3c6
	contractAddr2 = "0xf35bFE94C64d1dF4C86F0488ff687e307a2Ce3c6"
)

// get12ExecContract
func get12ExecContract() {
	//client, err := ethclient.Dial("<execution-layer-endpoint-url>")
	//wss://eth-sepolia.g.alchemy.com/v2/CtYhECjGkQZDMbZ1AkVvIRu9N8PyUX0Z
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/00fdc4fc92d945ea9e8b5a0157764345")
	//client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/CtYhECjGkQZDMbZ1AkVvIRu9N8PyUX0Z")
	if err != nil {
		log.Fatal(err)
	}
	storeContract, err := store.NewStore(common.HexToAddress(contractAddr2), client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("your private key")
	if err != nil {
		log.Fatal(err)
	}

	var key [32]byte
	var value [32]byte

	copy(key[:], []byte("demo_save_key"))
	copy(value[:], []byte("demo_save_value11111"))

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}
	tx, err := storeContract.SetItem(opt, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	callOpt := &bind.CallOpts{Context: context.Background()}
	valueInContract, err := storeContract.Items(callOpt, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("is value saving in contract equals to origin value:", valueInContract == value)
}

//API server listening at: 127.0.0.1:50258
//tx hash: 0x1ed7bbd58c8556de6775de48440e507499f61a32e489519896a00724d48bdc6e
//is value saving in contract equals to origin value: true
