package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {

	// 使用WebSocket协议连接到测试网节点
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个用于接收新区块头信息的通道
	headers := make(chan *types.Header)
	// 订阅新的区块头事件
	// 当区块链产生新区块时，区块头信息会被发送到headers通道
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		// 订阅失败 ，则记录错误并退出程序
		log.Fatal(err)
	}

	// 无限循环，持续监听来自订阅通道的消息
	for {
		select {
		case err := <-sub.Err(): // 从订阅的错误通道读取错误信息
			log.Fatal(err) // 如果订阅过程出现错误（如网络中断），则记录错误并退出
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time().Uint64())     // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7
		}
	}
}
