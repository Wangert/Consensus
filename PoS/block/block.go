package block

import (
	"time"
	"encoding/hex"
	"strconv"
	"crypto/sha256"
	"PoS/node"
	"math/rand"
)

type Block struct {
	//交易记录
	Data int
	//上一个区块Hash
	Prehash string
	//区块Hash
	Hash string
	//时间戳
	Timestamp string
	//区块高度
	Index int
	//挖矿节点
	Validator *node.Node  //记录挖矿的那个节点的地址
}

//创建创世区块
func GenesisBlock() Block {
	var genesisBlock = Block{0, "0", "", time.Now().String(), 0, nil}
	genesisBlock.Hash = hex.EncodeToString(calculateHash(&genesisBlock))

	return genesisBlock
}

//计算区块的Hash值
func calculateHash(block *Block) []byte {
	record := strconv.Itoa(block.Data) + strconv.Itoa(block.Index) + block.Prehash +
		block.Timestamp
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)

	return hashed
}

//PoS挖矿生成下一个区块
func GenerateNextBlock(oldBlock Block, data int, addr []*node.Node) Block {

	var newBlock Block
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Prehash = oldBlock.Hash
	newBlock.Data = data
	newBlock.Hash = hex.EncodeToString(calculateHash(&newBlock))

	//通过PoS计算由哪个村民挖矿
	//设计随机种子
	rand.Seed(time.Now().Unix())
	//产生[0,150)随机值
	var rd = rand.Intn(150)

	//选出挖矿的节点
	node := addr[rd]

	//设置当前区块挖矿人地址
	newBlock.Validator = node

	//当前node矿工的原有持币数增加
	node.Tokens = node.Tokens + 1


	return newBlock
}
