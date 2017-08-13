package simple_channel

import "fmt"

//测试用两份channel连接三个goroutine
func Channel_connection() {
	naturals := make(chan int)
	squares := make(chan int)
	//Counter
	go func() {
		for i := 0; ; i++ {
			naturals <- i
		}
	}()

	//Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	//Print the squares
	for {
		fmt.Println(<-squares)
		if j:=<-squares;j>10000{
			break
		}
	}

}

//测试并发打印10个hello world
func Channel_test() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go print(chs[i])
	}

	for _, ch := range chs {
		j := <-ch
		fmt.Println(j)
	}
}

func print(ch chan int) {
	fmt.Println("Hello world")
	ch <- 1
}
