package main

import (
	token "./contracts_erc20" // for demo
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func main() {
	// 建立连接：通过 ethclient.Dial连接到以太坊节点。这里使用的是 Cloudflare 提供的公共网关，无需 API 密钥，适合测试用途
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	// Golem (GNT) Address
	// // Golem (GNT) 代币的合约地址
	tokenAddress := common.HexToAddress("0xfadea654ea83c00e5003d2ea15c59830b65471c0")
	// 创建代币合约实例
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// 要查询余额的以太坊地址
	address := common.HexToAddress("0x25836239F7b632635F815689389C537133248edb")
	// 查询指定地址的代币余额（返回的是最小单位，如wei）
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	// 查询代币名称
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	// 查询代币符号（缩写）
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	// 查询代币的小数位数
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"
	fmt.Printf("wei: %s\n", bal)           // 原始余额（最小单位）         // "wei: 74605500647408739782407023"

	// 将余额转换为可读格式（考虑小数位数）
	fbal := new(big.Float)
	fbal.SetString(bal.String()) // 将big.Int转换为big.Float

	// 除以10^decimals来得到实际代币数量
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	// 格式化后的余额
	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}
