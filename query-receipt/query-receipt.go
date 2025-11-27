package main

import (
	"context"
	common2 "ethclient/common"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//代码首先通过 ethclient.Dial创建一个客户端实例，用于与以太坊节点通信
	client, err := ethclient.Dial(common2.ApiKey)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 代码通过指定区块号（例如 5671744）来查询完整的区块数据。
	// BlockByNumber方法返回的区块对象包含了该区块的所有元数据（如时间戳、难度值）以及包含的所有交易列表
	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	// 直接遍历区块交易列表：从 block.Transactions()返回的列表中直接获取交易对象
	for _, tx := range block.Transactions() {
		// 哈希
		fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		// 转账金额
		fmt.Println(tx.Value().String()) // 100000000000000000
		// Gas用量
		fmt.Println(tx.Gas()) // 21000
		// Gas价格
		fmt.Println(tx.GasPrice().Uint64()) // 100000000000
		fmt.Println(tx.Nonce())             // 245132
		fmt.Println(tx.Data())              // []
		fmt.Println(tx.To().Hex())          // 0x8F9aFd209339088Ced7Bc0f57Fe08566ADda3587

		// 恢复交易发送方：使用 types.Sender函数并传入基于当前链ID的EIP155签名器，
		// 可以从交易的签名中恢复出发送方的地址。这是验证交易来源的密码学安全方法
		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println("sender", sender.Hex()) // 0x2CdA41645F2dBffB852a605E92B185501801FC28
		} else {
			log.Fatal(err)
		}

		// 查询交易收据：通过 TransactionReceipt方法获取交易收据。收据包含了交易的执行结果，其中 Status字段为 1 表示交易成功，为 0 则表示失败。
		// Logs字段记录了交易执行过程中产生的事件日志，对于智能合约交互尤其重要
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status) // 1
		fmt.Println(receipt.Logs)   // []
		break
	}

	// 通过区块哈希和交易索引：先使用 TransactionCount获取区块内交易总数，再通过 TransactionInBlock根据索引位置获取特定交易
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("交易的笔数", count)
	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		break
	}

	// 直接通过交易哈希：使用 TransactionByHash可以直接查询任意交易，其返回的 isPending布尔值可以告诉你该交易是否还在等待被打包确认的状态
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isPending)
	fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5.Println(isPending)       // false
}
