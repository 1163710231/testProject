package main

import (
	"fmt"
	"sort"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param x int整型 采购单价
 * @param y int整型 采购单价
 * @param z int整型 采购单价
 * @return int整型
 */
func compare(x int, y int, z int) int {
	// write code here
	var nums = []int{x, y, z}
	sort.Ints(nums)
	return nums[2] - nums[0]
	//if x > y {
	//	if x > z { // x 最大
	//		if y > z { // z 最小
	//			return x - z
	//		} else { // y 最小
	//			return x - y
	//		}
	//	} else { // z 最大
	//		if x > y { // y 最小
	//			return z - y
	//		} else { // x 最小
	//			return z - x
	//		}
	//	}
	//} else {
	//	if y > z { // y 最大
	//		if x > z { // z 最小
	//			return y - z
	//		} else { // x 最小
	//			return y - x
	//		}
	//	} else { // z 最大
	//		if x > y { // y 最小
	//			return z - y
	//		} else { // x 最小
	//			return z - x
	//		}
	//	}
	//}
}

func main() {
	var x = 1
	var y = 2
	var z = 3
	fmt.Println(compare(x, y, z))
}
