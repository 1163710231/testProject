package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param material int整型一维数组 成品质量
 * @param standard int整型 质检标准
 * @return int整型一维数组
 */
func check(material []int, standard int) []int {
	// write code here
	result := make([]int, 0, len(material))
	for m := range material {
		if material[m] >= standard {
			result = append(result, material[m])
		}
	}
	return result
}

func main() {
	material := []int{2, 2, 3, 4, 6, 6, 3}
	fmt.Println(check(material, 3))
}
