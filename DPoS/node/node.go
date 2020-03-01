package node

import (
	"math/rand"
	"time"
)

//节点
type Node struct {
	//节点名称
	Name string
	//获得选票数量
	Votes int
}

//创建村民
func CreateNodes(n []*Node) {

	//设置随机种子
	rand.Seed(time.Now().Unix())

	node1 := Node{"node1", rand.Intn(10)}
	node2 := Node{"node2", rand.Intn(10)}
	node3 := Node{"node3", rand.Intn(10)}
	node4 := Node{"node4", rand.Intn(10)}
	node5 := Node{"node5", rand.Intn(10)}

	n[0] = &node1
	n[1] = &node2
	n[2] = &node3
	n[3] = &node4
	n[4] = &node5

}

//DPoS中选出票数最高的前n位
func SortNodes(n []*Node) []*Node {
	//对所有选民的票数进行排序
	for i := 0; i < 4; i++ {
		for j := 0; j < 4 - i; j++ {
			if n[j].Votes < n[j + 1].Votes {
				//二者的位置交换
				t := n[j]
				n[j] = n[j + 1]
				n[j + 1] = t
			}
		}
	}

	return n[:3]
}