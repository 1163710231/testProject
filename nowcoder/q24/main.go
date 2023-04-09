package main

import "fmt"

//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s1 int整型一维数组
 * @param s2 int整型一维数组
 * @return bool布尔型
 */
func equal(s1 []int, s2 []int) bool {
	// write code here
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{1, 2, 3, 4}
	fmt.Println(equal(s1, s2))
}
