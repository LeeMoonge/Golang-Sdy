package main

import "fmt"

//interface{}是万能数据类型
func myFunc(arg interface{})  {
	fmt.Println("arg =", arg)

	//interface{} 该如何区分 此时引用的底层属性到底是什么？
	//给interface{} 提供"类型断言"的机制
	value, check := arg.(string)  //类型断言
	if check {
		fmt.Printf("True! arg is string | value=%v\n", value)
	}else{
		fmt.Printf("False! arg not is string | value=%v | type=%T\n", value, value)
	}
}

type Pig struct {
	title string
}

func main()  {
	book := Pig{"Golang"}

	myFunc(book)
	myFunc(100)
	myFunc(3.14)
	myFunc("yqq")
}