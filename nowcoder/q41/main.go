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
func operate(a int, b int) []int {
	// write code here
	result := make([]int, 0, 5)
	result = append(result, a+b)
	result = append(result, a-b)
	result = append(result, a*b)
	result = append(result, a/b)
	result = append(result, a%b)
	return result
}

func main() {
	a, b := 2, 2
	fmt.Println(operate(a, b))
}
