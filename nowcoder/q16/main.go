package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param a int整型
 * @param b int整型
 * @return int整型一维数组
 */
func bitOperate(a int, b int) []int {
	// write code here
	var results = make([]int, 3)
	results[0] = a & b
	results[1] = a | b
	results[2] = a ^ b
	return results
}

func main() {
	fmt.Println(bitOperate(1, 2))
}
