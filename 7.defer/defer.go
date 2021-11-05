package main

//defer关键字：在函数结束之前执行的操作
//多个defer操作类似于压栈 最先入栈的最后执行，最后入栈的最先执行，依次类推

import "fmt"

func main()  {
	defer fmt.Println("in func ==> end1")
	defer fmt.Println("in func ==> end2")

	fmt.Println("in func ==> go 1")
	fmt.Println("in func ==> go 2")
}
