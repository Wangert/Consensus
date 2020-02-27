package block

import (
	"time"
	"encoding/hex"
	"strconv"
	"crypto/sha256"
	"fmt"
	"strings"
)

//创建区块的结构体，声明区块
//挖区块的方式，常用的有三种PoW,PoS,DPoS
type Block struct {

	//上一个区块的Hash值
	PreHash []byte
	//时间戳
	TimeStamp int64
	//交易信息
	Data []byte
	//当前区块链的Hash值
	Hash []byte
	//随机数
	Nonce int

}

//创建创世区块
func CreateGenesisBlock() *Block {
	var genesisBlock = &Block{[]byte{0}, time.Now().Unix(), []byte("wjt->jld 200BTC"), nil, 0}
	//计算当前区块的Hash值
	genesisBlock.getBlockHash()

	return genesisBlock
}

//计算block当前Hash值
func (block *Block) getBlockHash() string {
	var blockInfo = hex.EncodeToString(block.Data) + strconv.Itoa(block.Nonce) + string(block.TimeStamp) + hex.EncodeToString(block.PreHash)

	h := sha256.New()
	h.Write([]byte(blockInfo))
	hashed := h.Sum(nil)

	//设置计算出来的Hash值给当前block
	block.Hash = hashed

	return hex.EncodeToString(hashed)
}

//通过PoW挖矿的方式实现区块的产生
func GenerateNextBlock(oldBlock *Block) *Block {
	//假设难度系数为4，则挖的区块的Hash值前边必须有4个0，才算挖矿成功
	var newBlock = &Block{oldBlock.Hash, time.Now().Unix(), []byte("jld->wjt 100BTC"), nil, 0}

	//以步长为1增加Nonce值，最终使得当前区块Hash值前4位是0
	var nonce int = 1
	for {
		newBlock.Nonce = nonce
		hashString := newBlock.getBlockHash()

		fmt.Println("挖矿中：", hashString, "nonce:", newBlock.Nonce)

		if strings.HasPrefix(hashString, "0000") {
			fmt.Println("挖矿成功！")
			return newBlock
		}

		nonce++
	}

}