package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	n := runtime.GOMAXPROCS(4)
	fmt.Println("n=", n)

	fmt.Println("run in main goroutine")
	i := 1
	for k := 1; k < 1000000; k++ {
		go func() {
			j := 1
			for {
				j = j + 1
			}
		}()
		if i%10000 == 0 {
			fmt.Printf("%d goroutine started\n", i)
		}
		i++
	}

	time.Sleep(3600 * time.Second)
}
