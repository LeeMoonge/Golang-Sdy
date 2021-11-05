package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		defer fmt.Println("slave goroutine 结束 ...")

		c <- 666
	}()

	num := <- c
	fmt.Println("channel num =", num)
	fmt.Println("main goroutine 结束 ...")
}
