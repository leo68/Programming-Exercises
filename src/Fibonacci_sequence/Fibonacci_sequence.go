package Fibonacci_sequence

import (
	"fmt"
	"time"
)

//计算菲波那契数列的第n个元素值,n由用户输入
func Fibonacci_sequence() {
	fmt.Println("Please input the  sequence number of the fibonacci number you want to get:")
	var n int
	fmt.Scan(&n)
	go spinner(1000 * time.Millisecond)
	fibN := fib(n)
	fmt.Println("Fibonacci(%d)=%d\n", n, fibN)

}

//定义旋转器
func spinner(delay time.Duration) {
	for {
		for _, r := range `-/|\` {
			fmt.Print("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
