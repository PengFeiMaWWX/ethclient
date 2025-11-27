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

	// 使用ethclient.Dial通过一个API密钥（存储在common2.ApiKey中）建立与以太坊节点的连接
	client, err := ethclient.Dial(common2.ApiKey)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// 使用BlockByNumber方法查询指定区块号（这里是5671744）的完整区块信息。
	// 该方法返回一个区块对象，其中包含该区块的所有元数据和交易列表
	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	// 循环遍历集合并获取交易的信息。
	for _, tx := range block.Transactions() {
		// 交易hash 唯一标识
		fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		// 交易金额 以wei为单位的转账值。
		fmt.Println(tx.Value().String()) // 100000000000000000
		// Gas 相关数据 ， gas的消耗量和
		fmt.Println(tx.Gas()) // 21000
		// gas价格
		fmt.Println(tx.GasPrice().Uint64()) // 100000000000
		// 发送账户的交易序列号
		fmt.Println(tx.Nonce()) // 245132
		// 数据字段：通常用于智能合约交易
		fmt.Println(tx.Data()) // []
		// 接收方地址：交易接收者的以太坊地址
		fmt.Println(tx.To().Hex()) // 0x8F9aFd209339088Ced7Bc0f57Fe08566ADda3587

		// 获取交易发送方的地址
		// 使用types.Sender函数和EIP155签名器（types.NewEIP155Signer(chainID)）从交易中恢复发送方地址。
		// 这是通过验证交易签名来实现的，确保地址的准确性
		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println("sender", sender.Hex()) // 0x2CdA41645F2dBffB852a605E92B185501801FC28
		} else {
			log.Fatal(err)
		}

		// 查询交易收据
		// 通过TransactionReceipt方法获取交易的收据，其中包含交易执行结果（Status字段，1表示成功，0表示失败）和事件日志（Logs字段）
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status) // 1
		fmt.Println(receipt.Logs)   // []
		break
	}

	// 按区块哈希和索引查询：使用TransactionInBlock方法，通过区块哈希和交易在区块中的索引位置获取交易
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}
	// 循环遍历集合并获取交易的信息。
	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		break
	}

	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	// 直接按交易哈希查询：使用TransactionByHash方法直接查询交易，并返回交易是否待处理（isPending）
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(isPending)
	fmt.Println(tx.Hash().Hex()) // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5.Println(isPending)       // false
}
