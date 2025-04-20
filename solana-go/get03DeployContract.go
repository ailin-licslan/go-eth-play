package main

//import (
//	"context"
//	"github.com/blocto/solana-go-sdk/client"
//	"github.com/blocto/solana-go-sdk/rpc"
//	"os"
//)
//
//func deployTokenProgram() string {
//	c := client.NewClient(rpc.DevnetRPCEndpoint)
//
//	// 加载编译后的程序文件
//	programBin, _ := os.ReadFile("spl_token.so")
//
//	// 部署交易
//	txHash, _ := c.DeployProgramWithOpts(
//		context.Background(),
//		programBin,
//		rpc.DeployProgramOpts{
//			SkipPreflight:       false,
//			PreflightCommitment: rpc.CommitmentFinalized,
//		},
//	)
//
//	return txHash
//}
