package main

import "fmt"

//const 来定义枚举类型
const (
	//可以在const() 添加一个关键字 iota，每行的iota都会加1，第一行的iota默认值是0
	BEIJING = 10 * iota  //iota = 0
	SHANGHAI		//iota = 1
	SHENZHEN		//iota = 2
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 1, iota * 2
	i, j
)

func main () {
	//常量：只读属性
	const config int = 100

	fmt.Println("config =", config)

	//config = 200
	//# command-line-arguments
	//./demo3.go:11:9: cannot assign to config (declared const)
	//常量是不允许修改的

	fmt.Println("BEIJING =", BEIJING)
	fmt.Println("SHANGHAI =", SHANGHAI)
	fmt.Println("SHENZHEN =", SHENZHEN)

	fmt.Printf("a = %v, b = %v\n", a, b)
	fmt.Printf("c = %v, d = %v\n", c, d)
	fmt.Printf("e = %v, f = %v\n", e, f)
	fmt.Printf("g = %v, h = %v\n", g, h)
	fmt.Printf("i = %v, j = %v\n", i, j)
}
