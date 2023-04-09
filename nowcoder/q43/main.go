package main

import "fmt"

type student struct {
	name  string
	sex   bool
	age   int
	score int
	add   address
}

type address struct {
	province string
	city     string
}

func main() {
	s := student{
		name:  "小明",
		sex:   true,
		age:   23,
		score: 88,
		add: address{
			province: "湖南省",
			city:     "长沙市",
		},
	}

	fmt.Println(s.name)
	fmt.Println(s.sex)
	fmt.Println(s.age)
	fmt.Println(s.score)
	fmt.Println(s.add.province)
	fmt.Println(s.add.city)
}
