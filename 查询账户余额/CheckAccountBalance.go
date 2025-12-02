package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	// 要查询的以太坊地址（0x2583...8edb）
	account := common.HexToAddress("0x25836239F7b632635F815689389C537133248edb")
	// 查询该地址的最新余额（区块号设为nil表示获取最新区块的余额）
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err) // 查询失败则记录错误并退出
	}
	fmt.Println(balance) // 直接打印余额（以wei为单位）
	// 查询指定区块高度（5532993）时的账户余额
	blockNumber := big.NewInt(5532993)
	// 余额查询：BalanceAt方法接收三个参数：上下文、账户地址和区块号。当区块号设置为 nil时，返回的是最新确认区块的余额
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt) // 25729324269165216042

	// 将余额从wei转换为ETH（1 ETH = 10^18 wei）
	fbalance := new(big.Float)
	// 将big.Int转换为big.Float
	fbalance.SetString(balanceAt.String())
	// 除以10^18得到ETH单位
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041 人类可读的ETH余额
	// 查询待处理余额（包含已发送但尚未打包进区块的交易）
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance) // 25729324269165216042
}
