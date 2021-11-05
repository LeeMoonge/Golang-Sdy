package main

import "fmt"

//声明一种新的数据类型 myint，是int的别名
type myint int

//定义一个结构体
type Book struct {
	title string
	auth string
}

func changeBook(book *Book)  {
	//需使用指针传递才可以改变原有变量值，否则只是将book的副本拷贝进来
	book.auth = "小紫"
}

func main()  {
	var a myint = 666
	fmt.Printf("a=%v, type=%T\n", a, a)

	var book1 Book
	book1.title = "Golang"
	book1.auth = "yqq"

	fmt.Println("book1 =", book1)

	changeBook(&book1)

	fmt.Println("book1 =", book1)

}
