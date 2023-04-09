package main

import (
	"fmt"
	"math"
)

//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s int整型一维数组 评委给出分数
 * @return int整型一维数组
 */
func minAndMax(s []int) []int {
	// write code here
	var max = math.MinInt64
	var min = math.MaxInt64
	for i := range s {
		if s[i] > max {
			max = s[i]
		}
		if s[i] < min {
			min = s[i]
		}
	}
	var result []int
	result = []int{min, max}
	return result
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(minAndMax(s))
}
