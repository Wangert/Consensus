package main

import (
	"os"
	"fmt"
	"net/http"
	"io"
)

/*
拜占庭在Fabric中实现
 */

 
//声明节点信息，代表各个小国家
type nodeInfo struct {
	//节点标识
	id string
	//准备连接的地址:端口
	path string
	//服务器做出的响应
	writer http.ResponseWriter
}

//存放四个国家的地址
var nodeTable = make(map[string]string)


func main() {

 	//获取执行参数
 	userId := os.Args[1]
 	fmt.Println(userId)

 	//创建四个国家的地址
 	nodeTable = map[string]string {
 		"wjt":"localhost:1111",
 		"jld":"localhost:2222",
 		"mxl":"localhost:3333",
 		"xcx":"localhost:4444",
	}

	node := nodeInfo{userId, nodeTable[userId], nil}

	fmt.Println(node)

	//http协议的回调函数
	//http://localhost:1111/req?warTime=8888
	http.HandleFunc("/req", node.request)
	http.HandleFunc("/prePrepare", node.prePrepare)
	http.HandleFunc("/prepare", node.prepare)
	http.HandleFunc("/commit", node.commit)
	//启动服务器
	if err := http.ListenAndServe(node.path, nil); err != nil {
		fmt.Println(err)
	}
}

//此函数是http访问时req命令的请求回调函数,参数不能写反，因为系统调用会自动传参
func (node *nodeInfo)request(writer http.ResponseWriter, request *http.Request) {

	//设置允许解析参数
	request.ParseForm()
	fmt.Println(request.Form["warTime"][0])
	//如果有参数值，则继续处理
	if len(request.Form["warTime"]) > 0 {
		node.writer = writer
		//激活主节点，广播给其他节点，通过wjt节点向其他节点广播
		node.broadcast(request.Form["warTime"][0], "/prePrepare")

	}
}

//由主节点向其他节点做广播
func (node *nodeInfo)broadcast(msg string, path string) {
	for nodeId, url := range nodeTable {
		if nodeId == node.id {
			continue
		}

		//调用
		http.Get("http://" + url + path + "?warTime=" + msg + "&nodeId=" + node.id)
	}
}

func (node *nodeInfo)prePrepare(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	//fmt.Println("hello wjt")
	//再次做分发
	if len(request.Form["warTime"]) > 0 {
		//分发给其他人
		node.broadcast(request.Form["warTime"][0], "/prepare")
	}
}


func (node *nodeInfo)prepare(writer http.ResponseWriter, request *http.Request) {
	//接收参数
	request.ParseForm()
	//调用验证
	if len(request.Form["warTime"]) > 0 {
		fmt.Println(request.Form["warTime"][0])
	}

	if len(request.Form["nodeId"]) > 0 {
		fmt.Println(request.Form["nodeId"][0])
	}

	node.authentication(request)
}

var authenticationSuccess = true
var authenticationMap = make(map[string]string)

//获得除了本节点外的其他节点数据
func (node *nodeInfo)authentication(request *http.Request) {
	//接收参数
	request.ParseForm()

	if authenticationSuccess != false {
		if len(request.Form["nodeId"]) > 0 {
			authenticationMap[request.Form["nodeId"][0]] = "ok"
		}
	}

	//判断是否满足拜占庭
	if len(authenticationMap) > len(nodeTable) / 3 {
		//通过commit反馈给浏览器
		node.broadcast(request.Form["warTime"][0], "/commit")
	}
}

func (node *nodeInfo)commit(writer http.ResponseWriter, request *http.Request) {
	//给浏览器反馈响应
	//node.writer = writer
	if writer != nil {
		io.WriteString(node.writer, "ok")
	}
}