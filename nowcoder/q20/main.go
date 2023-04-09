package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param src int整型一维数组 源切片
 * @param des int整型一维数组 目的切片
 * @return int整型一维数组
 */
func sliceCopy(src []int, des []int) []int {
	// write code here
	des = make([]int, len(src), cap(src))
	copy(des, src)
	return des
}

func main() {
	var src = []int{1, 2}
	var des = make([]int, 0, 2)
	fmt.Println(sliceCopy(src, des))
}
