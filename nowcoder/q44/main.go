package main

import "fmt"

/**
	type 接口类型名 interface{
        方法名1( 参数列表1 ) 返回值列表1
        方法名2( 参数列表2 ) 返回值列表2
        …
    }
*/

type animal interface {
	sleep()
	eat()
}

type tiger struct {
}

func (t tiger) sleep() {
	fmt.Println("sleep")
}

func (t tiger) eat() {
	fmt.Println("eat")
}

func main() {
	var t animal
	t = tiger{}
	t.sleep()
	t.eat()
}
