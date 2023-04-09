package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param age int整型 年龄
 * @return string字符串
 */
func getAge(age int) string {
	// 婴儿(出生0-1岁)、幼儿(1-4岁)包含1岁、儿童(5-11)包含5岁、少年(12-18)包含12岁、青年(19-35)包含19岁、中年(36-59)包含36岁、老年(60以上)包含60岁
	if 0 <= age && age < 1 {
		return "婴儿"
	} else if 1 <= age && age < 5 {
		return "幼儿"
	} else if 5 <= age && age < 12 {
		return "儿童"
	} else if 12 <= age && age < 19 {
		return "少年"
	} else if 19 <= age && age < 36 {
		return "青年"
	} else if 36 <= age && age < 60 {
		return "中年"
	} else if 60 <= age {
		return "老年"
	} else {
		return "错误"
	}
}

func main() {
	age := 12
	fmt.Println(getAge(age))
}
