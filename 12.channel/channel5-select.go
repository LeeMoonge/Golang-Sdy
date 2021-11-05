package main

import "fmt"

func Fibonacii(c chan int, quit chan int)  {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			//如果c可写，则该case就会进来
			x = y
			y = x + y
		case <-quit:
			//如果quit可读，则该case会进来
			fmt.Println("quit ...")
			return
		}
	}
}


func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i ++{
			fmt.Println("channel c =", <-c)
		}
		quit <- 0
	}()

	//main Go程执行
	Fibonacii(c, quit)
}
