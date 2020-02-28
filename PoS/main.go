package main

import (
	"fmt"
	"PoS/node"
	"PoS/block"
	"strconv"
)

/*
现实PoS挖矿的原理
 */


//创建5个村民
var n = make([]node.Node, 5)
//存放每个村民的地址
var addr = make([]*node.Node, 150)



func main() {

	//初始化区块链节点
	node.InitNode(n, addr)

	//创建创世区块
	var genesisBlock = block.GenesisBlock()

	//创建新区块
	var newBlock = block.GenerateNextBlock(genesisBlock, 39, addr)

	fmt.Println("Validator's address：" + newBlock.Validator.Address)
	fmt.Println("Validator's token：" + strconv.Itoa(newBlock.Validator.Tokens))

	fmt.Println("=========================================================")
	fmt.Println("Information of the new block:")
	fmt.Println(newBlock)

}
