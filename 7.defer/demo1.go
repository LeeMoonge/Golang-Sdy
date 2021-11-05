package main

import "fmt"

/*
	知识点一：defer的执行顺序
	多个defer操作类似于压栈 最先入栈的最后执行，最后入栈的最先执行，依次类推
*/

func demo1()  {
	fmt.Println("demo1 echo A")
}

func demo2()  {
	fmt.Println("demo2 echo B")
}

func demo3()  {
	fmt.Println("demo3 echo C")
}

func main()  {
	defer demo1()
	defer demo2()
	defer demo3()
}