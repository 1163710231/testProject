package main

import "fmt"

type networkDelayError struct {
}

func (nde networkDelayError) Error() string {
	return "网络延迟"
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param ping int整型 网络延迟值
 * @return string字符串
 */
func defineerr(ping int) string {
	// write code here
	if ping > 100 {
		nde := networkDelayError{}
		return nde.Error()
	} else {
		return ""
	}
}

func main() {
	ping := 1213
	fmt.Println(defineerr(ping))
}
