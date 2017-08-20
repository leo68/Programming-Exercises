package SortingAlgorithm

import "fmt"

type SortInterface interface {
	sort()
}

type Sortor struct {
	name string
}

func Run() {
	array := []int{6, 1, 2, 5, 7, 9, 0, 3, 5, 6, 44, 76554, 34, 13, 6, 345, 4, 5}
	fmt.Println("待排数组:", array)
	learnsort := Sortor{name: "快速排序"}
	learnsort.Quicksort(array)
	fmt.Print(learnsort.name+": ", array)
}
