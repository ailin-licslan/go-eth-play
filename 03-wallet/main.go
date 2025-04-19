package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	//generate private key
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	pData := crypto.FromECDSA(pvk)
	fmt.Println(hexutil.Encode(pData))

	//generate public key
	puData := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println(hexutil.Encode(puData))

	//generate public address
	fmt.Println(crypto.PubkeyToAddress(pvk.PublicKey).Hex())
}
