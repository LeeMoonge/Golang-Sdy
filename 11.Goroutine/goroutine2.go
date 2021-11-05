package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {


	//用go创建一个形参为空 返回值为空的一个匿名函数
	go func() {
		defer fmt.Println("defer A")

		func() {
			defer fmt.Println("defer B")
			//退出当前的goroutine
			runtime.Goexit()
			fmt.Println("print B")
		}()

		fmt.Println("print A")
		return
	}()

	//在main函数中创建一个带形参的匿名函数
	func(a int, b int) bool {
		fmt.Println("a =", a, "b =", b)
		return true
	}(666, 777)


	for  {
		time.Sleep(1 * time.Second)
	}
}