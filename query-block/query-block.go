package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
)

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	// 连接以太坊节点
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/<API_KEY>")
	if err != nil {
		log.Fatal(err)
	}

	// 设置目标区块号
	blockNumber := big.NewInt(5671744)

	// 获取区块头信息
	// 可以理解为区块的元数据或者摘要， 包含了重要信息， 但是不包含交易信息
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	// 区块高度
	fmt.Println(header.Number.Uint64()) // 5671744
	// 区块挖出的时间戳
	fmt.Println(header.Time) // 1712798400
	// 工作量证明（pow）难度值，Sepolia 通常为0
	fmt.Println(header.Difficulty.Uint64()) // 0
	// 区块哈希 区块的唯一标识
	fmt.Println(header.Hash().Hex()) // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5

	if err != nil {
		log.Fatal(err)
	}

	// 完整的区块信息(包含了所有的交易信息)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Time())                // 1712798400
	fmt.Println(block.Difficulty().Uint64()) // 0
	fmt.Println(block.Hash().Hex())          // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	// 获取区块中的交易数量
	fmt.Println(len(block.Transactions())) // 70
	// 通过区块哈希验证交易数量
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 70
}
