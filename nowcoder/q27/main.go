package main

import "fmt"

func main() {
	// map[小明:60 小王:70 张三:95 张伟:88 李四:98 王五:100]
	m := make(map[string]int, 6)
	m["小明"] = 60
	m["小王"] = 70
	m["张三"] = 95
	m["张伟"] = 88
	m["李四"] = 98
	m["王五"] = 100
	fmt.Println(m)
}
