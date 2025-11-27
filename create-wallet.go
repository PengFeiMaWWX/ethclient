package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	// 创建新的钱包需要 ，提供生成随机私钥的GenerateKey方法
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	// 生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// FromECDSA 方法将其转换为字节。
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// 包将它转换为十六进制字符串，该包提供了一个带有字节切片的 Encode 方法。 然后我们在十六进制编码之后删除'0x'
	// 私钥被作为密码， 用于签署交易 ， 拥有私钥就可以访问钱包账户
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // 去掉'0x' 私钥的hash字符串 5f1d4364483c44dadcf7291f33e4edd2b98470b365abb1eba2dc71e41c74798e

	// 通过私钥派生公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	// 剥离了 0x 和前 2 个字符 04，它是 EC 前缀，不是必需的
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("from pubKey:", hexutil.Encode(publicKeyBytes)[4:]) // 去掉'0x04' 2a1c0666624daa4a413a64bd7174cffaeed3f7eb032bd10cda465bedc7a5d242fd7cd8ac578bcf3f93d0f8472f3c842ef50ea6c0c64e9f2759376eb7334f61eb

	// 现在我们拥有公钥，就可以轻松生成你经常看到的公共地址。
	// 为了做到这一点，go-ethereum 加密包有一个 PubkeyToAddress 方法，它接受一个 ECDSA 公钥，并返回公共地址。
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address) // 0xC234cB0C200EF493F6B28D2e4c9958243F0FE4F6

	// 公共地址就是公钥的 Keccak-256 哈希 然后我们取最后 40 个字符（20 个字节）并用“0x”作为前缀
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:])) // 完整哈希 0xb458af8cff99d26064d3da8fc234cb0c200ef493f6b28d2e4c9958243f0fe4f6
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))        // 原长32位，截去12位，保留后20位 0xc234cb0c200ef493f6b28d2e4c9958243f0fe4f6
}
