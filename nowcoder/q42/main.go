package main

import "fmt"

type student struct {
	name  string
	sex   bool
	age   int
	score int
}

func main() {
	s := student{"小明", true, 23, 88}

	fmt.Println(s.name)
	fmt.Println(s.sex)
	fmt.Println(s.age)
	fmt.Println(s.score)
}
