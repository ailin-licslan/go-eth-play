//以下是使用 Solana Go SDK 查询第一个区块数据的示例代码：

package main

import (
	"context"

	"log"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/system"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
)

// FUarP2p5EnxD66vVDL4PWRoWMzA56ZVHG24hpEDFShEz
var feePayer, _ = types.AccountFromBase58("4TMFNY9ntAn3CHzguSAvDNLPRoQTaK3sWbQQXdDXaE6KWRBLufGL6PJdsD2koiEe3gGmMdRK3aAw7sikGNksHJrN")

// 9aE476sH92Vz7DMPyq5WLPkrKWivxeuTKEFKd2sZZcde
var alice, _ = types.AccountFromBase58("4voSPg3tYuWbKzimpQK9EbXHmuyy5fUrtXvpLDMLkmY6TRncaTHAKGD8jUg3maB5Jbrd9CkQg4qjJMyN6sQvnEF2")

func main() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// to fetch recent blockhash
	recentBlockhashResponse, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf("failed to get recent blockhash, err: %v", err)
	}

	// create a transfer tx
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{feePayer, alice},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: recentBlockhashResponse.Blockhash,
			Instructions: []types.Instruction{
				system.Transfer(system.TransferParam{
					From:   alice.PublicKey,
					To:     common.PublicKeyFromString("2xNweLHLqrbx4zo1waDvgWJHgsUpPj8Y8icbAFeR4a8i"),
					Amount: 1e8, // 0.1 SOL
				}),
			},
		}),
	})
	if err != nil {
		log.Fatalf("failed to new a transaction, err: %v", err)
	}

	// send tx
	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("failed to send tx, err: %v", err)
	}

	log.Println("txhash:", txhash)
}

//// 代币转账事件结构
//type TokenTransfer struct {
//	Source      common.PublicKey `json:"source"`
//	Destination common.PublicKey `json:"destination"`
//	Amount      uint64           `json:"amount"`
//}
//
//func parseTokenTransfer(logs []string) (*TokenTransfer, error) {
//	// 查找转账指令日志
//	for _, l := range logs {
//		if token.IsTransferInstructionLog(l) {
//			// 示例日志格式："Program log: Instruction: Transfer,
//			// Amount=100000000"
//			return &TokenTransfer{
//				Source:      common.PublicKeyFromString(logs[2]), // 账户位置固定
//				Destination: common.PublicKeyFromString(logs[3]),
//				Amount:      extractAmountFromLog(l),
//			}, nil
//		}
//	}
//	return nil, errors.New("未找到转账事件")
//}
//
//// 从日志字符串提取金额
//func extractAmountFromLog(log string) uint64 {
//	var amount struct {
//		Value uint64 `json:"Amount"`
//	}
//	json.Unmarshal([]byte(log[15:]), &amount) // 跳过"Program log: "前缀
//	return amount.Value
//}
