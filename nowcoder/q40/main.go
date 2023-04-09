package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param x int整型
 * @return int整型
 */
func absfunc(x int) int {
	// write code here
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func main() {
	x := -12
	fmt.Println(absfunc(x))
}
