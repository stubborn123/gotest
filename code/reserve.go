package main

import "fmt"

//实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。
//就是利用这个切片，二分法
func main() {
	fmt.Printf(reverStr("123456"))
}

func reverStr(s string) string {
	//就是把字符串转换为切片[]rune(s)
	str := []rune(s)
	len := len(str)

	if len > 5000 {
		return s
	}

	for i := 0; i < len/2; i++ {
		//利用go可以多个值赋值（叫做“多重赋值”）的情况 1,2 = a,b
		str[i], str[len-1-i] = str[len-1-i], str[i]
	}
	return string(str)
}
