package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

func getWallet() {

	//首先生成一个新的钱包，我们需要导入 go-ethereum crypto 包，该包提供用于生成随机私钥的 GenerateKey 方法。
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	//如果已经有了私钥的 Hex 字符串，也可以使用 HexToECDSA 方法恢复私钥：
	//privateKey, err := crypto.HexToECDSA("ccec5314acec3d18eae81b6bd988b844fc4f7f7d3c828b351de6d0fede02d3f2")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//然后我们可以通过导入 golang crypto/ecdsa 包并使用 FromECDSA 方法将其转换为字节。
	privateKeyBytes := crypto.FromECDSA(privateKey)

	//我们现在可以使用 go-ethereum hexutil 包将它转换为十六进制字符串，该包提供了一个带有字节切片的 Encode 方法。 然后我们在十六进制编码之后删除“0x”。
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // 去掉'0x'

	//这就是用于签署交易的私钥，将被视为密码，永远不应该被共享给别人，因为谁拥有它可以访问你的所有资产。由于公钥是从私钥派生的，因此 go-ethereum 的加密私钥具有一个返回公钥的 Public 方法。
	publicKey := privateKey.Public()

	//将其转换为十六进制的过程与我们使用转化私钥的过程类似。 我们剥离了 0x 和前 2 个字符 04，它是 EC 前缀，不是必需的。
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("from pubKey:", hexutil.Encode(publicKeyBytes)[4:]) // 去掉'0x04'

	//现在我们拥有公钥，就可以轻松生成你经常看到的公共地址。 为了做到这一点，go-ethereum 加密包有一个 PubkeyToAddress 方法，它接受一个 ECDSA 公钥，并返回公共地址。
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)

	//公共地址其实就是公钥的 Keccak-256 哈希，然后我们取最后 40 个字符（20 个字节）并用“0x”作为前缀。 以下是使用 golang.org/x/crypto/sha3 的 Keccak256 函数手动完成的方法。
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 原长32位，截去12位，保留后20位
}

//以下是打印结果
//214384e772ee960de78c2d11b4a362648a2cfeb51c1754884fe4d573828f1114
//from pubKey: fd20d8e6bb2487c0455f5ad585b70fafe461acb15c23bf87276d451ff5391cc154ceaf77526763c63a4da2c687f66450c7eb33db5ac992855cf6d2b5106d9372
//0x1FB969716DC77e9f7B055e6d15F21756492b3ff0
//full: 0x871e5017d7c5272e0edcc2801fb969716dc77e9f7b055e6d15f21756492b3ff0
//0x1fb969716dc77e9f7b055e6d15f21756492b3ff0
