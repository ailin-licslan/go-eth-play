package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

// get9SubBlock
func get9SubBlock() {
	//https://sepolia.infura.io/v3/00fdc4fc92d945ea9e8b5a0157764345
	//wss://sepolia.infura.io/ws/v3/00fdc4fc92d945ea9e8b5a0157764345
	//client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	//wss://eth-sepolia.g.alchemy.com/v2/CtYhECjGkQZDMbZ1AkVvIRu9N8PyUX0Z
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/CtYhECjGkQZDMbZ1AkVvIRu9N8PyUX0Z")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			//fmt.Println(header.Hash().Hex())
			//fmt.Println(header.Number.Uint64())
			//fmt.Println(header.Time)
			//fmt.Println(header.Nonce.Uint64())

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}
	}
}

//API server listening at: 127.0.0.1:51007
//0xd76d3cf606aeb069334831333ee744444922afadb5aa9e6be0da9affd6efe308
//0xd76d3cf606aeb069334831333ee744444922afadb5aa9e6be0da9affd6efe308
//8141708
//1744951308
//0
//112
//0x0210a71ed87b0d18fb539ff612aaaf063507215d0457c1efe30490d55473b782
//0x0210a71ed87b0d18fb539ff612aaaf063507215d0457c1efe30490d55473b782
//8141709
//1744951320
//0
//116
//0xebfe44688a36756ba892e366cf83b8f48b149f96eb70a868f7df4c6aa2070299
//0xebfe44688a36756ba892e366cf83b8f48b149f96eb70a868f7df4c6aa2070299
//8141710
//1744951332
//0
//130
//0x1651f53b34bbc552747cfddaa50f1f46b8875360ff283d88fd5a3c3fbdedea0d
