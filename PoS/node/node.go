package node

//创建全节点类型，可以理解成持有币的村民类型
type Node struct {
	//持有币的个数
	Tokens int
	//持有币的时间
	Days int
	//地址
	Address string
}


//初始化节点
func InitNode(n []Node, addr []*Node) {
	n[0] = Node{10, 1, "11111"}
	n[1] = Node{20, 1, "22222"}
	n[2] = Node{30, 1, "33333"}
	n[3] = Node{40, 1, "44444"}
	n[4] = Node{50, 1, "55555"}

	cnt := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < n[i].Tokens * n[i].Days; j++ {
			addr[cnt] = &n[i]

			cnt++
		}
	}
}