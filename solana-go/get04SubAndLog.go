package main

//
//import (
//	"context"
//	"fmt"
//	"log"
//
//	"github.com/blocto/solana-go-sdk/client"
//	"github.com/blocto/solana-go-sdk/common"
//	"github.com/blocto/solana-go-sdk/program/system"
//	"github.com/blocto/solana-go-sdk/rpc"
//	"github.com/blocto/solana-go-sdk/types"
//
//	"github.com/portto/solana-go-sdk/rpc/ws"
//)
//
//func main() {
//	// 创建WebSocket客户端（连接到DevNet）
//	wsClient, err := ws.Connect(context.Background(), rpc.DevnetWS)
//	if err != nil {
//		panic("连接失败: " + err.Error())
//	}
//
//	// 订阅程序日志（示例：SPL Token程序）
//	sub, err := wsClient.LogsSubscribeMatching(
//		"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",
//		rpc.CommitmentConfirmed,
//	)
//	if err != nil {
//		panic("订阅失败: " + err.Error())
//	}
//	defer sub.Unsubscribe()
//
//	// 实时处理事件
//	for {
//		select {
//		case log := <-sub.Channel():
//			fmt.Printf("[程序日志] 账户: %s\n", log.Value.Pubkey)
//			fmt.Printf("      日志内容: %+v\n", log.Value.Logs)
//		case err := <-sub.Err():
//			fmt.Printf("监听错误: %v\n", err)
//			return
//		}
//	}
//}

//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//import (
//"encoding/json"
//"github.com/portto/solana-go-sdk/common"
//"github.com/portto/solana-go-sdk/program/token"
//)
//
//// 代币转账事件结构
//type TokenTransfer struct {
//	Source     common.PublicKey `json:"source"`
//	Destination common.PublicKey `json:"destination"`
//	Amount     uint64           `json:"amount"`
//}
//
//func parseTokenTransfer(logs []string) (*TokenTransfer, error) {
//	// 查找转账指令日志
//	for _, l := range logs {
//		if token.IsTransferInstructionLog(l) {
//			// 示例日志格式："Program log: Instruction: Transfer,
//			// Amount=100000000"
//			return &TokenTransfer{
//				Source:     common.PublicKeyFromString(logs[2]), // 账户位置固定
//				Destination: common.PublicKeyFromString(logs[3]),
//				Amount:     extractAmountFromLog(l),
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
