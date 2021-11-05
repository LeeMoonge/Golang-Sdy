package main

import (
	"fmt"
	"time"
)

//子协程 子goroutine
func myTask() {
	i := 0
	for {
		i ++
		fmt.Println("myTask Goroutine run i=", i)
		time.Sleep(1 * time.Second)
	}
}

//主goroutine
func main() {
	//创建一个goroutine go程 去执行myTask()流程
	go myTask()

	//如果main程序结束，所有创建的子协程都会退出
	fmt.Println("main Goroutine exit ...")

	//mainI := 0
	//for {
	//	mainI ++
	//	fmt.Println("main Goroutine run mainI=", mainI)
	//	time.Sleep(1 * time.Second)
	//}
}
