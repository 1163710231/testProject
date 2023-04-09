package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param des string字符串
 * @param src string字符串
 * @return bool布尔型
 */
func canConstruct(des string, src string) bool {
	// write code here
	m := make(map[byte]int)
	ld, ls := len(des), len(src)
	for i := 0; i < ls; i++ {
		_, exist := m[src[i]]
		if exist {
			m[src[i]] += 1
		} else {
			m[src[i]] = 1
		}
	}

	for i := 0; i < ld; i++ {
		charRemain, exist := m[des[i]]
		if !exist || charRemain == 0 {
			return false
		} else {
			m[des[i]] -= 1
		}
	}
	return true
}

func main() {
	des := "abcd"
	src := "dbca"
	fmt.Println(canConstruct(des, src))
}
