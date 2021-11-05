package main

import "fmt"

//go 中变量都包含一个pair对，

func main()  {

	var a string
	//pair<static type:string, value: "aceld">
	a = "aceld"

	//pair<type: string, value: "aceld">
	var allType interface{}
	allType = a

	str, check := allType.(string)
	fmt.Println(str, check)
}