package main

import (
	"DPoS/block"
	"DPoS/node"
	"fmt"

)

/*
实现DPoS原理
 */

//创建数组，保存区块链
var Blockchain = make([]*block.Block, 0)

//创建数组，保存所有节点
var n = make([]*node.Node, 5)


func main() {

	//创建节点
	node.CreateNodes(n)

	//选定指定人员为矿工，该实例选3名票数高者
	threeMiners := node.SortNodes(n)

	//记录创世区块
	var genesisBlock *block.Block

	//获取更新后区块链和创世区块
	Blockchain, genesisBlock = block.GenesisBlock(Blockchain)


	//记录新区块
	var newBlock *block.Block

	//被选中的节点轮流做矿工
	//创建新区块
	Blockchain, newBlock = block.GenerateNextBlock(Blockchain, *genesisBlock, 39)
	//设置选取的第一位为矿工
	newBlock.SetDelegate(threeMiners[0])

	Blockchain, newBlock = block.GenerateNextBlock(Blockchain, *newBlock, 10)
	newBlock.SetDelegate(threeMiners[1])

	Blockchain, newBlock = block.GenerateNextBlock(Blockchain, *newBlock, 16)
	newBlock.SetDelegate(threeMiners[2])

	//循环遍历出区块链内容
	for _, blockInfo := range Blockchain {
		fmt.Println(blockInfo.Delegate.Name, blockInfo.Hash, blockInfo.Data)
	}

}
