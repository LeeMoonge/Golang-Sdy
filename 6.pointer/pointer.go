package main

import "fmt"

func changevalue(a , b int)  {
	temp := a
	a = b
	b = temp
}

func changevalue2(a *int, b *int)  {
	temp := *a
	*a = *b
	*b = temp
}

func main()  {
	var a int = 10
	var b int = 20

	changevalue(a, b)
	fmt.Println("a =", a, "b =", b)

	changevalue2(&a, &b)
	fmt.Println("a =", a, "b =", b)

	var p *int  //一级指针
	p = &a
	//p := &a
	fmt.Println("address=> a=", &a, "p=", p)

	var pp **int  //二级指针
	pp = &p
	fmt.Println("address=> pp=", pp, "&p=", &p)

}
