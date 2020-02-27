package main

import (
	"PoW/block"
	"encoding/hex"
	"fmt"
)

/*
完成比特币中PoW挖矿的共识算法
 */


func main() {
	//创建创世区块
	var genesisBlock = block.CreateGenesisBlock()
	//fmt.Println(genesisBlock.Hash)

	var newBlock = block.GenerateNextBlock(genesisBlock)
	fmt.Println("挖矿成功:", hex.EncodeToString(newBlock.Hash), "nonce:", newBlock.Nonce)

}