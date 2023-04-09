package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param array int整型二维数组
 * @return int整型一维数组
 */
func transform(array [][]int) []int {
	// write code here
	length := len(array) * len(array[0])
	result := make([]int, 0, length)
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[0]); j++ {
			result = append(result, array[i][j])
		}
	}
	return result
}

func main() {
	array := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	result := transform(array)
	fmt.Println(result)
}
