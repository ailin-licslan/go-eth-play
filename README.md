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
├── 2-获取区块信息  get01ChainInfo.go 
├── 3-获取交易信息  get02TradeInfo.go
├── 4-获取收据信息  get03ReceiptsInfo.go
├── 5-钱包生成     get04Wallet.go  
├── 6-转ETH       get05ETHTransfer.go
├── 7-转token     get06TokenTransfer.go
├── 8-账户余额     get08TokenBalance.go
├── 9-订阅区块     get09SubBlock.go
├── 10-部署合约    get10DeployContract.go
├── 11-加载合约    get11LoadContract.go
├── 12-执行合约    get12ExecContract.go
├── 13-合约事件    get13EventContract.go
```


