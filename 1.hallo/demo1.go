package main //程序包名

import (
	"fmt"
	"time"
)

/*
import "time"
import "fmt"
*/

//main函数
func main() {  //函数的 { 一定是和函数名在同一行，否则编译错误
	//golang的表达式，加 ";" 和不加都可以
	fmt.Println("malaohu!!!")

	time.Sleep(1 * time.Second)
}
