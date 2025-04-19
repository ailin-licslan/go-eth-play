package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// 加密钱包使用
func main() {
	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "passwordd"
	// a, err := key.NewAccount(password)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(a.Address)

	b, err := ioutil.ReadFile("./wallet/UTC--2025-04-19T05-32-18.099797400Z--f4d7f07a9a8303625eb1e3587a4c6fb0d2ce70e0")
	if err != nil {
		log.Fatal(err)
	}
	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal(err)
	}
	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("Priv", hexutil.Encode(pData))

	pData = crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Pub", hexutil.Encode(pData))

	fmt.Println("Add", crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())
}
