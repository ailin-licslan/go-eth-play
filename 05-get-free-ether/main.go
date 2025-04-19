package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var (
	infuraURL = "https://sepolia.infura.io/v3/00fdc4fc92d945ea9e8b5a0157764xxx"
	url       = "https://kovan.infura.io/v3/0c7b3f204f37416388610fb274b0452c"
	murl      = "https://mainnet.infura.io/v3/0c7b3f204f37416388610fb274b0452c"
)

func main() {

	//生成2个钱包  key store type wallet
	//ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	//_, err := ks.NewAccount("password")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//_, err = ks.NewAccount("password")
	//if err != nil {
	//	log.Fatal(err)
	//}
	// "f4d7f07a9a8303625eb1e3587a4c6fb0d2ce70e0"
	// "619b219598ff921fea69658c261b685fdb5cc4a2"

	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	a1 := common.HexToAddress("f4d7f07a9a8303625eb1e3587a4c6fb0d2ce70e0")
	a2 := common.HexToAddress("619b219598ff921fea69658c261b685fdb5cc4a2")

	b1, err := client.BalanceAt(context.Background(), a1, nil)
	if err != nil {
		log.Fatal(err)
	}

	b2, err := client.BalanceAt(context.Background(), a2, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance 1:", b1)
	fmt.Println("Balance 2:", b2)
}
