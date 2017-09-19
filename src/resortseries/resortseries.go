/*
[编程题] 重排数列
时间限制：1秒
空间限制：100768K
小易有一个长度为N的正整数数列A = {A[1], A[2], A[3]..., A[N]}。
牛博士给小易出了一个难题:
对数列A进行重新排列,使数列A满足所有的A[i] * A[i + 1](1 ≤ i ≤ N - 1)都是4的倍数。
小易现在需要判断一个数列是否可以重排之后满足牛博士的要求。
输入描述:
输入的第一行为数列的个数t(1 ≤ t ≤ 10),
接下来每两行描述一个数列A,第一行为数列长度n(1 ≤ n ≤ 10^5)
第二行为n个正整数A[i](1 ≤ A[i] ≤ 10^9)


输出描述:
对于每个数列输出一行表示是否可以满足牛博士要求,如果可以输出Yes,否则输出No。

输入例子1:
2
3
1 10 100
4
1 2 3 4

输出例子1:
Yes
No
*/

//运行超时:您的程序未能在规定时间内运行结束，请检查是否循环有错或算法复杂度过大。
//case通过率为80.00%
package resortseries

import "fmt"

func ResortSeries() {
	var t int
	fmt.Scanf("%d", &t)
	var slice [][]int

	for i := 0; t > i; i++ {
		var t1 int
		fmt.Scanf("%d", &t1)
		var tmpArr []int
		for t1 > 0 {
			var t2 int
			fmt.Scanf("%d", &t2)
			tmpArr = append(tmpArr, t2)
			t1--
		}
		slice = append(slice, tmpArr)
	}

	for i, _ := range slice {
		j := len(slice[i])
		m := 0
		total := 0
		num := 0
		for ; m < j; m++ {
			if slice[i][m]%4 == 0 {
				total++
				continue
			}
			if slice[i][m]%2 == 0 {
				num++
			}
		}
		if num == 0 {
			if j-total <= total+1 {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		} else {
			if j-total-num+1 <= total+1 {
				fmt.Println("Yes")
			} else {
				fmt.Println("No	")
			}
		}
	}
}
