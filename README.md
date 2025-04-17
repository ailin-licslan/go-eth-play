# GO-ETH-PLAY


## 使用指南
### 下载
```bash
git clone https://github.com/ailin-licslin/go-eth-play.git
```
### 主要内容（如何与链上交互）
1. go mod projectName  && go build 编译后执行 可以debugger调试
```sql
注册Infura账号 获取到API_KEY
client, err := ethclient.Dial("https://sepolia.infura.io/v3/<API_KEY>"

go get  github.com/ethereum/go-ethereum/common
go get	github.com/ethereum/go-ethereum/core/types
go get	github.com/ethereum/go-ethereum/crypto
go get	github.com/ethereum/go-ethereum/ethclient

./go-ethereum.exe
```

2. 文件分层(go-ethereum learn)  
```
├── 1-eth-transfer         Let's get strated!
│   └── 查询链上相关数据
│   └── 钱包相互转账操作
├── 2-获取区块信息  get1ChainInfo.go 
├── 3-获取交易信息  get2TradeInfo.go
├── 4-获取收据信息  get3ReceiptsInfo.go
├── 5-钱包生成     get4Wallet.go  
├── 6-账号转ETH    get5ETHTransfer.go
│   └── pointer.go
```


