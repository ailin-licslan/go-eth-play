package store

//Windows env:
//1.Please make sure you already have installed npm on your computer
//2.npm install -g solc
//3.solcjs --bin Store.sol    solcjs --abi Store.sol
//4.go install github.com/ethereum/go-ethereum/cmd/abigen@latest
//5.执行 cd /c/Users/Administrator/go/pkg/mod/github.com/ethereum/go-ethereum@v1.15.8/cmd/abigen
//6.执行 go build 来生成 abigen.exe
//7.执行 ./abigen --bin=Store_sol_Store.bin --abi=Store_sol_Store.abi --pkg=store --out=store.go

//Linux/mac env:
//1.Please make sure you already have installed npm on your computer
//2.npm install -g solc
//3.solcjs --bin Store.sol    solcjs --abi Store.sol
//4.go get -u github.com/ethereum/go-ethereum/cmd/abigen
//5.执行 cd /c/put your directory here/go/pkg/mod/github.com/ethereum/go-ethereum@v1.15.8/xxx/abigen
//6.执行 make && make devtools 来生成 abigen
//7.执行 ./abigen --bin=Store_sol_Store.bin --abi=Store_sol_Store.abi --pkg=store --out=store.go
