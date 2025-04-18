package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

// getETHTransfer
func get5ETHTransfer() {

	//转账交易包括打算转账的以太币数量，燃气限额，燃气价格，一个自增数(nonce)，接收地址以及可选择性的添加的数据。
	//在发送到以太坊网络之前，必须使用发送方的私钥对该交易进行签名。

	//NETWORK : sepolia.infura.io  这里请填一个注册好的地址  可以参考目录 go-eth-play/eth-transfer/InfuraRPC.png
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/<API_KEY>")
	if err != nil {
		log.Fatal(err)
	}

	//0xBC9c5bD5eC8f4FE7Dd0988EC236931122dc69f79                        metamask address  [dev] address
	//ab640f1cebd3854b116911111111111111111111111111111111111111111111  metamask address  [dev] privateKey
	//fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19  这里可以从metamask钱包导出一个测试钱包账号的私钥 不要用生产账号!!!
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//下一步是设置我们将要转移的 ETH 数量。 但是我们必须将 ETH 以太转换为 wei，因为这是以太坊区块链所使用的。
	//以太网支持最多 18 个小数位，因此 1 个 ETH 为 1 加 18 个零。
	//value := big.NewInt(1000000000000000) // in wei (1 eth)  0.001  //tx sent: 0x1a13f6d0e2aa02aea366ebd02483025a21586c1f66c26447b2ccaab532f9340c
	//value := big.NewInt(1000000000000000) // in wei (1 eth)  0.001  //tx sent: 0x5337f15bf9398df06f5e71dedcdee14ac79cd3656aef658987c9ff4a0e3e31cd
	value := big.NewInt(5000000000000000) //0.005  tx sent: 0xa250d9b97376c3f18d14e8f74fff1b9519e97ce80be8298b6a88587c1afb7af3
	//ETH 转账的燃气应设上限为“21000”单位。
	gasLimit := uint64(21000) // in units

	//然而，燃气价格总是根据市场需求和用户愿意支付的价格而波动的，因此对燃气价格进行硬编码有时并不理想。
	//go-ethereum 客户端提供 SuggestGasPrice 函数，用于根据'x'个先前块来获得平均燃气价格。
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//0xBC9c5bD5eC8f4FE7Dd0988EC236931122dc69f79  dev address (from address)
	//0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	//0xa9E54ea6745cf2ec23235710eF4D38D7d3985bd2  developing address  (to address)
	//接下来我们弄清楚我们将 ETH 发送给谁。
	toAddress := common.HexToAddress("0xa9E54ea6745cf2ec23235710eF4D38D7d3985bd2") //metamask address here  找一个测试账号的钱包地址
	//toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte

	//现在我们最终可以通过导入 go-ethereum core/types 包并调用 NewTransaction 来生成我们的未签名以太坊事务，这个函数需要接收 nonce，
	//地址，值，燃气上限值，燃气价格和可选发的数据。 发送 ETH 的数据字段为“nil”。 在与智能合约进行交互时，我们将使用数据字段，仅仅转账以太币是不需要数据字段的。
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	//下一步是使用发件人的私钥对事务进行签名。 为此，我们调用 SignTx 方法，该方法接受一个未签名的事务和我们之前构造的私钥。
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//SignTx 方法需要 EIP155 签名者，这个也需要我们先从客户端拿到链 ID。
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//现在通过在 client 实例调用 SendTransaction 来将已签名的事务广播到整个网络。
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
