package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s string字符串一维数组
 * @return string字符串
 */
func join(s []string) string {
	// write code here
	var result string
	for _, v := range s {
		result += v
	}
	return result
}

func main() {
	var s = []string{"aaa", "bbb", "ccc"}
	//s := []string{"aaa", "bbb", "ccc"}
	var r string
	r = join(s)
	fmt.Println(r)
}
