package main

import (
	"fmt"
	//Fibonacci "github.com/Programming-Exercises/src/Fibonacci_sequence"
	//timeServer "github.com/Programming-Exercises/src/time_server"
	//timeClent "github.com/Programming-Exercises/src/time_client"
	simpleChannel "github.com/Programming-Exercises/src/simple_channel"

)

func main() {
	fmt.Println("Hello world!")
	//计算菲波那契数列
	//Fibonacci.Fibonacci_sequence()
	//时间服务器，每秒钟刷新一次，需要特定客户端接收
	//go timeServer.TimeServer()
	//时间客户端，通过本地8000端口获取时间，按秒变化
	//timeClent.Netcat()
	//测试并发打印10个Hello world
	//simpleChannel.Channel_test()
	//测试用两个channel连接三个goroutine,打印自然数的平方(小于10000)
	simpleChannel.Channel_connection()
}


