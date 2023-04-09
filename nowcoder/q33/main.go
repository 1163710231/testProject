package main

import "fmt"

//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param hight double浮点型 身高
 * @return bool布尔型
 */
func ispay(hight float64) bool {
	// write code here
	if hight-160.0 < 0 {
		return false
	} else {
		return true
	}
}

func main() {
	height := 163.3445
	fmt.Println(ispay(height))
}
