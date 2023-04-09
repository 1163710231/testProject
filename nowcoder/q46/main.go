package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param t double浮点型 体温
 * @return string字符串
 */
func temperature(t float64) (result string) {
	// write code here
	result = ""
	defer func() {
		err := recover()
		if err != nil {
			result = "体温异常"
		}
	}()
	if t > 37.5 {
		panic("体温异常，弹出警告！")
	}
	return result
}

func main() {
	t := 38.9
	fmt.Println(temperature(t))
}
