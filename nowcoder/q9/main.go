package main

import (
	"fmt"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 将一个正整数转换为字符串
 * @param num int整型 给定的正整数
 * @return string字符串
 */
func formatstr(num int) string {
	// write code here
	//return strconv.Itoa(num)
	var s = fmt.Sprintf("%d", num)
	return s
}

func main() {
	var num = 123123
	fmt.Println(formatstr(num))
}
