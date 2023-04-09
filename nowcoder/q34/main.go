package main

import "fmt"

//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param forwards string字符串 推箱子方向
 * @return bool布尔型
 */
func pushBox(forwards string) bool {
	// write code here
	U, D, L, R := 0, 0, 0, 0
	for i := range forwards {
		if forwards[i] == 'U' {
			U++
		} else if forwards[i] == 'D' {
			D++
		} else if forwards[i] == 'L' {
			L++
		} else if forwards[i] == 'R' {
			R++
		} else {
			return false
		}
	}

	return U == D && L == R
}

func main() {
	forwards := "UUDDLRLR"
	fmt.Println(pushBox(forwards))
}
