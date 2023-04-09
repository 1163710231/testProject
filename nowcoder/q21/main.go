package main

import "fmt"

func main() {
	var n float64 = 365
	var m float64 = 365
	var x float64 = 1

	for i := 0; i < 40; i++ {
		x *= n / m
		n--
	}

	//n--
	//for i := 0; i < 40; i++ {
	//	x *= n / m
	//}

	fmt.Println(x)
	fmt.Println(1 - x)
}
