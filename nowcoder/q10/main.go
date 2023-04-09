package main

import (
	"fmt"
	"strconv"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param a string字符串
 * @param b string字符串
 * @return string字符串
 */
func sum(a string, b string) string {
	// write code here
	//var ai, _ = strconv.Atoi(a)
	//var bi, _ = strconv.Atoi(b)
	var ai, _ = strconv.ParseInt(a, 10, 64)
	var bi, _ = strconv.ParseInt(b, 10, 64)
	var sum = ai + bi
	return strconv.Itoa(int(sum))
}

func main() {
	var a = "12"
	var b = "34"
	fmt.Println(sum(a, b))
}
