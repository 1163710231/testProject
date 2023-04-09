package main

import (
	"fmt"
	"sort"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s int整型一维数组
 * @return int整型一维数组
 */
func getNoRepeat(s []int) []int {
	// write code here
	m := make(map[int]bool)
	for i := range s {
		_, exist := m[s[i]]
		if exist {
			m[s[i]] = false
		} else {
			m[s[i]] = true
		}
	}

	var result []int
	for key, value := range m {
		if value {
			result = append(result, key)
		}
	}
	sort.Ints(result)
	return result
}

func main() {
	s := []int{3, 3, 2, 2, 5, 5, 1, 2, 2}
	fmt.Println(getNoRepeat(s))
}
