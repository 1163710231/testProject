package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param f double浮点型 华氏温度
 * @return double浮点型
 */
func temperature(f float64) float64 {
	// write code here
	return 5 * (f - 32.0) / 9
}

func main() {
	var f = 100.0
	fmt.Println(temperature(f))
}
