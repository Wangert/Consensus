package block

import (
	"time"
	"strconv"
	"crypto/sha256"
	"encoding/hex"
	"DPoS/node"
)

//区块
type Block struct {
	//区块高度
	Index int
	//时间戳
	Timestamp string
	//前区块Hash
	Prehash string
	//区块Hash
	Hash string
	//数据
	Data int
	//增加代理
	Delegate *node.Node
}

//产生新的区块
func GenerateNextBlock(Blockchain []*Block, oldBlock Block, data int) ([]*Block, *Block) {
	var newBlock = Block{oldBlock.Index + 1, time.Now().String(), oldBlock.Hash, "", data, nil}

	calculateHash(&newBlock)
	//添加到区块链中
	Blockchain = append(Blockchain, &newBlock)
	return Blockchain, &newBlock
}

//计算Hash
func calculateHash(block *Block) {
	record := strconv.Itoa(block.Index) + strconv.Itoa(block.Data) + block.Timestamp + block.Prehash

	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	block.Hash = hex.EncodeToString(hashed)

}

//设置代理人
func (block *Block) SetDelegate(node *node.Node) {
	block.Delegate = node
}

//创建创世区块
func GenesisBlock(Blockchain []*Block) ([]*Block, *Block) {
	var genesisBlock = Block{0, time.Now().String(), "0", "", 0, &node.Node{"", 0}}
	calculateHash(&genesisBlock)
	Blockchain = append(Blockchain, &genesisBlock)

	return Blockchain, &genesisBlock
}