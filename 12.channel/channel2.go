package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)  //创建带有缓存区的channel

	fmt.Println("len(c) =", len(c), "cap(c) =", cap(c))

	go func() {
		for i := 0; i < 4; i++{
			c <- i
			fmt.Println("Go子协程正在执行，发送到channel的元素为:", i, "len(c) =", len(c), "cap(c) =", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++{
		num := <-c
		fmt.Println("num =", num)
	}
}