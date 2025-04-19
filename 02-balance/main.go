package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://sepolia.infura.io/v3/00fdc4fc92d945ea9e8b5a0157764xxx"
var ganacheURL = "http://localhost:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), infuraURL)
	if err != nil {
		log.Fatalf("Error to create a ether client:%v", err)
	}
	defer client.Close()

	//last one block
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Error to get a block:%v", err)
	}
	fmt.Println("The block number:", block.Number())
	addr := "0xBC9c5bD5eC8f4FE7Dd0988EC236931122dc69f79"
	//addr := "0x944E2E2c632C4D6aF195DC3Bdec9C17F6fc6F600"
	address := common.HexToAddress(addr)

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Error to get the balance:%v", err)
	}
	fmt.Println("The balance:", balance)

	// 1 ether = 10^18 wei
	fBlance := new(big.Float)
	fBlance.SetString(balance.String())
	fmt.Println(fBlance)
	// 10*10*10*10*...18
	blanceEther := new(big.Float).Quo(fBlance, big.NewFloat(math.Pow10(18)))
	fmt.Println(blanceEther)
}
