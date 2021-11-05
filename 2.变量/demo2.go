package main

/*
	四种变量的声明方式
*/

import (
	"fmt"
)

//声明全局变量只能使用 方法一、方法二、方法三
var gA string
var gB int = 666
var gC = 3.14

//用方法四声明全局变量会报错，该语法只能用于函数体内部
//gD := "dddd"

func main()  {
	//方法一：声明一个变量，默认的值是0
	var a int
	var b string
	var c float64
	fmt.Println("a = ", a, "b = ", b, "c = ", c)

	//方法二：声明一个变量，初始化一个值
	var aa int = 100
	fmt.Println("aa = ", aa)
	fmt.Printf("type of aa %T\n", aa)

	//方法三：在初始化的时候，可以省去声明数据类型，通过值自动匹配当前的变量的数据类型
	var aaa = 100
	var bbb = "abcd"
	var ccc = 3.14
	fmt.Println("aaa =", aaa, "bbb =", bbb, "ccc =", ccc)
	fmt.Printf( "type(aaa=%v)=%T, type(bbb=%s)=%T, type(ccc=%v)=%T\n", aaa, aaa, bbb, bbb, ccc, ccc)

	//方法四：(常用的方法)省去var关键字，直接自动匹配
	aaaa := 100
	fmt.Println("aaaa =", aaaa)
	fmt.Printf("type(aaaa=%v)=%T\n", aaaa, aaaa)

	fmt.Println("=========全局变量==========")

	fmt.Println("gA =", gA, "gB =", gB, "gC =", gC)
	fmt.Printf("type(gA=%v)=%T, type(gB=%v)=%T, type(gC=%v)=%T", gA, gA, gB, gB, gC, gC)

	//===================声明多个变量======================
	//方法一：
	var xx, yy int = 111, 222
	fmt.Println(xx, yy)
	var ww, zz = 333, "aoe"
	fmt.Println(ww, zz)

	//方法二：
	var (
		xxx int = 999
		fff bool = true
	)
	fmt.Println(xxx, fff)

}
