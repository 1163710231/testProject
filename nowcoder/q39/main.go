package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param i int整型 数字
 * @return int整型
 */
func factorial(i int) int {
	// write code here
	if i == 0 || i == 1 {
		return 1
	} else {
		return i * factorial(i-1)
	}
}

func main() {
	i := 4
	fmt.Println(factorial(i))
}
