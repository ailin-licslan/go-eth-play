package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	token "go-ethereum/abi"
	"log"
	"math"
	"math/big"
)

// get8TokenBalance
func get8TokenBalance() {

	//Wallet address:   0xBC9c5bD5eC8f4FE7Dd0988EC236931122dc69f79 (deploy address)
	//contract address: 0x1295afec813450401f84335e6f83ae4cbfe89e43 (token address)
	//to address:       0xa9E54ea6745cf2ec23235710eF4D38D7d3985bd2 (to address)
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/00fdc4fc92d945ea9e8b5a0157764xxx")
	//client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com") //免费地址
	if err != nil {
		log.Fatal(err)
	}
	// Golem (GNT) Address
	tokenAddress := common.HexToAddress("0xfadea654ea83c00e5003d2ea15c59830b65471c0")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("0x25836239F7b632635F815689389C537133248edb")
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
	fmt.Printf("wei: %s\n", bal)           // "wei: 74605500647408739782407023"
	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}

//API server listening at: 127.0.0.1:57533
//name: RCCDemoToken
//symbol: RDT
//decimals: 18
//wei: 9997000000000000000000000
//balance: 9997000.000000
