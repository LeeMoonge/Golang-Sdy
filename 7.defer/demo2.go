package main

import "fmt"

/*
	知识点二：defer和return谁先谁后
			return之后的语句先执行，defer之后的语句后执行，defer关键字是在函数执行完毕后再执行的
*/

func deferFunc() int {
	fmt.Println("deferFunc execult")
	return 0
}

func returnFunc() int {
	fmt.Println("returnFunc execult")
	return 0
}

func returnAndDefer() int {

	defer deferFunc()

	return returnFunc()
}


func main()  {
	returnAndDefer()
}











//======================================================================================================================
//func help() func(a int) int {
//	num := 10
//	return func(a int) int {
//		num += a
//		return num
//	}
//}
//
//func main()  {
//	fun := help()
//	//fmt.Println(fun(1))
//	fmt.Println("fun() =", fun(1))
//}