package main

import "fmt"

//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s int整型一维数组
 * @return int整型一维数组
 */
func convert(s []int) []int {
	// write code here
	var l = len(s)
	for i := 0; i < l>>1; i++ {
		s[i], s[l-i-1] = s[l-i-1], s[i]
	}
	return s
}

func main() {
	s := []int{1, 2, 3, 4}
	s = convert(s)
	fmt.Println(s)
}
