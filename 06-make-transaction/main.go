package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	url  = "https://sepolia.infura.io/v3/00fdc4fc92d945ea9e8b5a0157764345" //测试网
	murl = "https://mainnet.infura.io/v3/00fdc4fc92d945ea9e8b5a0157764345" //主网
)

// transferring Ether between accounts using Golang
func main() {

	//创建2个钱包
	// ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	// _, err := ks.NewAccount("password")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = ks.NewAccount("password")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// "1f7ecea2fa83cc4a7de969f11d16a40edf9023d7"
	// "1e41ca1ccfc06597525c966a986b35a09e22358d"

	//连接eth test network
	client, err := ethclient.Dial(url)
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

	//打印钱包余额
	fmt.Println("Balance 1:", b1)
	fmt.Println("Balance 2:", b2)

	//get nonce
	nonce, err := client.PendingNonceAt(context.Background(), a1)
	if err != nil {
		log.Fatal(err)
	}

	//转多少 1 ether = 1000000000000000000 wei
	amount := big.NewInt(100000000000000000)

	//getPrice
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//创建交易信息   a2 是接收方地址
	tx := types.NewTransaction(nonce, a2, amount, 21000, gasPrice, nil)

	//chainId
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//获取加密钱包信息  最终为了获取私钥
	b, err := ioutil.ReadFile("wallet/UTC--2025-04-19T05-32-18.099797400Z--f4d7f07a9a8303625eb1e3587a4c6fb0d2ce70e0")
	if err != nil {
		log.Fatal(err)
	}

	//获取private key
	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal(err)
	}

	//签名交易信息
	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), key.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	//广播交易到区块链网络
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	//打印交易hash
	fmt.Printf("tx sent: %s", tx.Hash().Hex())
}
