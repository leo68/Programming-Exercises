package simple_channel

import "fmt"

//test1:测试用两份channel连接三个goroutine
func Channel_connection() {
	naturals := make(chan int)
	squares := make(chan int)
	//Counter
	go func() {
		for i := 0; i <= 100; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	//Squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	//Print the squares
	for x := range squares {
		fmt.Println(x)
		/*
			if j:=<-squares;j>10000{
				break
			}
		*/
	}

}

//test2:测试并发打印10个hello world
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

//test3:以下为新的实例，单方向channel
func counter(out chan<- int) {
	for i := 0; i <= 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for j := range in {
		out <- j * j
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func OnewayChannel() {
	naturals := make(chan int)
	squares := make(chan int)
	//以下调用将包含隐式类型转换，chan int转化为chan<- int，注意此过程不可逆
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

//test4:带缓存的channels，尾部插入，头部读取
//关键函数cap,len
func MirroredQuery() {
	response := make(chan string, 3)
	go func() { response <- request("asia.gopl.io") }()
	go func() { response <- request("europe.gopl.io") }()
	go func() { response <- request("americas.gopl.io") }()
	//若只读取一个则会造成goroutine泄露，不会自动回收，BUG
	fmt.Println(<-response)
}

func request(hostname string) (response string) {
	response = hostname
	return response
}
