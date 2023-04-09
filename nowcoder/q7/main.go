package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s string字符串
 * @return int整型
 */
func count(s string) int {
	// write code here
	var ss = []rune(s)
	//ss := []rune(s)
	return len(ss)
}
func main() {
	var s = "小明的英文名叫jack"
	var r = count(s)
	fmt.Println(r)
}
