/*
[编程题] 字符串碎片
时间限制：1秒
空间限制：32768K
一个由小写字母组成的字符串可以看成一些同一字母的最大碎片组成的。例如,"aaabbaaac"是由下面碎片组成的:'aaa','bb','c'。牛牛现在给定一个字符串,请你帮助计算这个字符串的所有碎片的平均长度是多少。

输入描述:
输入包括一个字符串s,字符串s的长度length(1 ≤ length ≤ 50),s只含小写字母('a'-'z')


输出描述:
输出一个整数,表示所有碎片的平均长度,四舍五入保留两位小数。

如样例所示: s = "aaabbaaac"
所有碎片的平均长度 = (3 + 2 + 3 + 1) / 4 = 2.25

输入例子1:
aaabbaaac

输出例子1:
2.25
 */


package stringpiece

import "fmt"

func StringPiecen() {
	var str string
	fmt.Scanf("%s", &str)
	if len(str)<0||len(str)>50{
		fmt.Println("Input error.")
		return
	}
	result:= carculateStringPiecen(str)
	s := fmt.Sprintf("%.2f", result)
	fmt.Println(s)
}

func carculateStringPiecen(str string)float32{
	var i,j int
	j=1
	for i<len(str)-1{
		if str[i:i+1]==str[i+1:i+2]{
			i++
			continue
		}
		i++
		j++

	}
	return float32(len(str))/float32(j)
}
