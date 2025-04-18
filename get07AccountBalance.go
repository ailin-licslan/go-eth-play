package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

// get7AccountBalance
func get7AccountBalance() {

	//Wallet address:   0xBC9c5bD5eC8f4FE7Dd0988EC236931122dc69f79 (deploy address)
	//contract address: 0x1295afec813450401f84335e6f83ae4cbfe89e43 (token address)
	//to address:       0xa9E54ea6745cf2ec23235710eF4D38D7d3985bd2 (to address)
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/00fdc4fc92d945ea9e8b5a0157764xxx")
	//client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com") //免费地址

	//client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	account := common.HexToAddress("0xa9E54ea6745cf2ec23235710eF4D38D7d3985bd2")
	//account := common.HexToAddress("0xBC9c5bD5eC8f4FE7Dd0988EC236931122dc69f79")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // balance = 218299403785650151

	//balance 转 ETH unit
	fmt.Println(getETHUnit(balance))

	blockNumber := big.NewInt(8141034)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt) // 25729324269165216042
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance) // 25729324269165216042
}

func getETHUnit(a *big.Int) *big.Float {

	fbalance := new(big.Float)
	fbalance.SetString(a.String())
	result := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	// 格式化输出结果，保留5位小数
	fmt.Printf("%.5f\n", result)

	return result
}
