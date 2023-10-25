package main

import (
	"bufio"
	//	"net/rpc"
	"flag"
	"net/rpc"
	"os"

	"fmt"
	//	"bufio"
	//	"os"
	"uk.ac.bris.cs/distributed2/secretstrings/stubs"
)

func main() {
	server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)
	//TODO: connect to the RPC server and send the request(s)
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()

	file, _ := os.Open("wordlist")
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		a := reader.Text() //分别读取file中的text,string的message
		fmt.Println("Called: ", a)
		request := stubs.Request{Message: a} //请求包含需要改动信息
		response := new(stubs.Response)      //反转后的信息
		client.Call(stubs.PremiumReverseHandler, request, response)
		fmt.Println("Responded: ", response.Message)
	}
	/*request := stubs.Request{Message: "Hello"}
	response := new(stubs.Response)
	client.Call(stubs.PremiumReverseHandler, request, response) //通过rpc在client调用server函数实现字符倒转
	fmt.Println("Responded: " + response.Message)*/
}
