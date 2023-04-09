package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {

			if i == j {
				fmt.Printf("%d*%d=%-3d\n", j, i, i*j)
			} else {
				fmt.Printf("%d*%d=%-3d", j, i, i*j)
			}
		}
	}
}
