package main

import (
	"fmt"
	"strconv"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param x int整型
 * @return bool布尔型
 */
func isPalindrome(x int) bool {
	// write code here
	var s = strconv.Itoa(x)
	var l = len(s)
	for i := 0; i < l>>1; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var x = 12421
	fmt.Println(isPalindrome(x))
}
