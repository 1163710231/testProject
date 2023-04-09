package main

import "fmt"

func main() {
	// [帽子 围巾 衣服 裤子 袜子]
	s := make([]string, 0, 5)
	s = append(s, "帽子", "围巾", "衣服", "裤子", "袜子")
	fmt.Println(s)
}
