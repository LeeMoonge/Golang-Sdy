package main

import "fmt"

func main() {
	c := make(chan int, 5)  //必须先定义一个channel，对一个nil channel无论收发都会阻塞

	go func() {
		for i := 0; i < 5; i++{
			c <- i
			fmt.Println("子Go程向channel中发送元素：", i)
		}
		//close 可以关闭一个channel
		close(c)  //确定没有发送的元素建议关闭channel，不然会引发程序报死锁错误
	}()

	for {
		//该语法先执行if后面的表达式，再判断ok是否符合条件
		//ok如果为True表示channel为开启状态，如果是false表示channel为关闭状态
		if data, ok := <-c; ok {
			fmt.Println("main Go程从channel从channel获取的元素：", data)
		}else {
			fmt.Println("channel已关闭！")
			break
		}
	}

	fmt.Println("main Go程结束 。。。")
}
