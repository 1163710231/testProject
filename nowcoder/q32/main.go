package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param score int整型 分数
 * @return string字符串
 */
func judgeScore(score int) string {
	// write code here
	switch {
	case 0 <= score && score < 60:
		return "不及格"
	case 60 <= score && score < 80:
		return "中等"
	case 80 <= score && score < 90:
		return "良好"
	case 90 <= score && score <= 100:
		return "优秀"
	default:
		return "错误"
	}
}

func main() {
	score := 23
	fmt.Println(judgeScore(score))
}
