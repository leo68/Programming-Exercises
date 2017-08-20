package worldLength

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

/*
题目描述
计算字符串最后一个单词的长度，单词以空格隔开。
输入描述:
一行字符串，非空，长度小于5000。
输出描述:
整数N，最后一个单词的长度。
示例1
输入
hello world
输出
5
 */
func WorldLength() {
	inputReader := bufio.NewReader(os.Stdin)
	str, _ := inputReader.ReadString('\n')
	strSlice:=strings.Fields(str)
	len1 := len(strSlice)
	lastworld := strSlice[len1-1]
	fmt.Print(len(lastworld))
}

func Str_contain() {
	inputReader := bufio.NewReader(os.Stdin)
	str, _ := inputReader.ReadString('\n')
	ch, _ := inputReader.ReadString('\n')
	strSli:=strings.Split(ch,"")
	strSli1:=strings.ToLower(strSli[0])
	strSli2:=strings.ToUpper(strSli[0])
	num1:=strings.Count(str,strSli1)
	num2:=strings.Count(str,strSli2)
	fmt.Print(num1+num2)
}
