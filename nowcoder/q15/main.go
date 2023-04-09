package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param a bool布尔型
 * @param b bool布尔型
 * @return bool布尔型一维数组
 */
func logicalOperation(a bool, b bool) []bool {
	// write code here
	//var results = make([]bool, 4)
	//results[0] = a && b
	//results[1] = a || b
	//results[2] = !a
	//results[3] = !b
	var results = make([]bool, 0)
	results = append(results, a && b)
	results = append(results, a || b)
	results = append(results, !a)
	results = append(results, !b)
	return results
}

func main() {
	fmt.Println(logicalOperation(true, false))
}
