package SortingAlgorithm

import (
	"fmt"
	"strconv"
)

/*
算法步骤：
设要排序的数组是A[0]……A[N-1]，首先任意选取一个数据（通常选用数组的第一个数）作为关键数据，然后将所有比它小的数都放到它前面，
所有比它大的数都放到它后面，这个过程称为一趟快速排序。值得注意的是，快速排序不是一种稳定的排序算法，也就是说，多个相同的值的相对位置也许会在算法结束时产生变动。
一趟快速排序的算法是：
1）设置两个变量i、j，排序开始的时候：i=0，j=N-1；
2）以第一个数组元素作为关键数据，赋值给key，即key=A[0]；
3）从j开始向前搜索，即由后开始向前搜索(j–)，找到第一个小于key的值A[j]，将A[j]和A[i]互换；
4）从i开始向后搜索，即由前开始向后搜索(i++)，找到第一个大于key的A[i]，将A[i]和A[j]互换；
5）重复第3、4步，直到i=j； (3,4步中，没找到符合条件的值，即3中A[j]不小于key,4中A[i]不大于key的时候改变j、i的值，
使得j=j-1，i=i+1，直至找到为止。找到符合条件的值，进行交换的时候i， j指针位置不变。另外，i==j这一过程一定正好是i+或j-完成的时候，此时令循环结束）。

快速排序是不稳定的。最理想情况算法时间复杂度O(nlog2n)，最坏O(n ^2)。
*/

func (s Sortor) Quicksort(array []int) {
	if len(array) <= 1 {
		return
	}
	mid := array[0]
	i := 1
	head, tail := 0, len(array)-1
	for head < tail {
		if array[i] > mid {
			array[i], array[tail] = array[tail], array[i]
			tail--
		} else {
			array[i], array[head] = array[head], array[i]
			head++
			i++
		}
	}
	array[head] = mid
	s.Quicksort(array[:head])
	s.Quicksort(array[head+1:])
}

/*
	题目描述
明明想在学校中请一些同学一起做一项问卷调查，为了实验的客观性，他先用计算机生成了N个1到1000之间的随机整数（N≤1000），对于其中重复的数字，只保留一个，把其余相同的数去掉，不同的数对应着不同的学生的学号。然后再把这些数从小到大排序，按照排好的顺序去找同学做调查。请你协助明明完成“去重”与“排序”的工作。


Input Param
     n               输入随机数的个数
 inputArray      n个随机整数组成的数组

Return Value
     OutputArray    输出处理后的随机整数


注：测试用例保证输入参数的正确性，答题者无需验证。测试用例不止一组。



输入描述:
输入多行，先输入随机整数的个数，再输入相应个数的整数
输出描述:
返回多行，处理后的结果
示例1
输入

11
10
20
40
32
67
40
20
89
300
400
15
输出

10
15
20
32
40
67
89
300
400
 */
func RemoveSortImp() {
	//var input []int
	input := make([]int, 1000)
	fmt.Scan(&input[0])
	//fmt.Print(input)
	for i := 1; i <= input[0]; i++ {
		fmt.Scan(&input[i])
	}
	i := input[0]
	array := input[1 : i+1]
	//fmt.Print(array)
	RemoveSort(array)
	array1 := RemoveDuplicates(array)
	//for i:=0;i<len(array1);i++{
	//	fmt.Println(array1[i])
	//}
	for mm := range array1 {
		//fmt.Println(array1[mm])
		fmt.Print(array1[mm], "\n")
	}
}

func RemoveSort(array []int) {
	if len(array) <= 1 {
		return
	}
	mid := array[0]
	i := 1
	head, tail := 0, len(array)-1
	for head < tail {
		if array[i] > mid {
			array[i], array[tail] = array[tail], array[i]
			tail--
		} else {
			array[i], array[head] = array[head], array[i]
			head++
			i++
		}
	}
	array[head] = mid
	RemoveSort(array[:head])
	RemoveSort(array[head+1:])
}

func RemoveDuplicates(a []int) (ret []int) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if i > 0 && a[i-1] == a[i] {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

/*
题目描述
•连续输入字符串，请按长度为8拆分每个字符串后输出到新的字符串数组；
•长度不是8整数倍的字符串请在后面补数字0，空字符串不处理。
输入描述:
连续输入字符串(输入2次,每个字符串长度小于100)
输出描述:
输出到长度为8的新字符串数组
示例1
输入

abc
123456789
输出

abc00000
12345678
90000000
 */
func LengthFormat() {
	input := make([]string, 2)
	fmt.Scan(&input[0])
	fmt.Scan(&input[1])

	var output []string
	m := len(input)
	for j := 0; j < m; j++ {
		length := len(input[j])
		for i := 0; i < length/8; i++ {
			output = append(output, input[j][(i*8):(i*8+8)])
		}
		if length/8*8 != length {
			m := length - length/8*8
			a := input[j][(length / 8 * 8):length]
			for ; 8-m > 0; m++ {
				a = a + "0"
			}
			output = append(output, a)
		}
	}

	for m := range output {
		fmt.Print(output[m], "\n")
	}

}

/*
写出一个程序，接受一个十六进制的数值字符串，输出该数值的十进制字符串。（多组同时输入 ）
输入描述:
输入一个十六进制的数值字符串。
输出描述:
输出该数值的十进制字符串。
示例1
输入

0xA
输出

10
 */
func StrTransform(){
	var a string
	fmt.Scan(&a)
	if a[0:2]!="0x"{
		return
	}
	b:=a[2:]
	var num uint64=0
	for i:=0;i<len(b);i++{
		base,_:=strconv.ParseInt(b[i:i+1],16,10)
		num=num*16+uint64(base)
	}
	fmt.Print(num)
}