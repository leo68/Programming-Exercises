package main

import (
	"fmt"
	//Fibonacci "github.com/Programming-Exercises/src/Fibonacci_sequence"
	timeServer "github.com/Programming-Exercises/src/time_server"
)

func main(){
	fmt.Println("Hello world!")
	//计算菲波那契数列
	//Fibonacci.Fibonacci_sequence()
	//时间服务器，每秒钟刷新一次，需要特定客户端接收
	timeServer.TimeServer()

}
