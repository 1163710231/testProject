package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param s string字符串
 * @return char字符型
 */
func character(s string) byte {
	// write code here
	m := make(map[byte]int)
	for i := range s {
		_, exist := m[s[i]]
		if exist {
			m[s[i]] += 1
		} else {
			m[s[i]] = 1
		}
	}

	max := -1
	var maxChar byte
	for key, value := range m {
		if value > max {
			maxChar = key
			max = value
		}
	}
	return maxChar
}

func main() {
	s := "yyds"
	fmt.Println(character(s))
}
