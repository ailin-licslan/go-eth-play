//以下是使用 Solana Go SDK 查询第一个区块数据的示例代码：

package main

import (
	"context"
	"fmt"
	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
)

// go get github.com/portto/solana-go-sdk@v1.25.0
func main() {

	// 创建RPC客户端（连接到DevNet）
	newClient := client.NewClient(rpc.DevnetRPCEndpoint)
	ctx := context.Background()

	// 查询第一个区块（区块高度为 0）的数据
	block, err := newClient.GetBlock(ctx, 0)
	if err != nil {
		fmt.Printf("获取区块数据失败: %v\n", err)
		return
	}

	fmt.Printf("第一个区块的哈希: %s\n", block.Blockhash)
	fmt.Printf("第一个区块的高度: %d\n", block.BlockHeight)
	fmt.Printf("第一个区块的交易数量: %d\n", len(block.Transactions))
}

//上述代码中：
//
//1.首先创建了一个连接到 Solana 测试网的 RPC 客户端。
//2.然后使用 client.GetBlock 方法来获取高度为 0 的第一个区块的数据。
//3.最后将获取到的区块的哈希、高度和交易数量打印输出。
//
//请注意，这里使用的是测试网的 RPC 端点 rpc.TestNet_RPCURL ，如果要在主网或其他网络上操作，你需要相应地更改 RPC 端点地址。
