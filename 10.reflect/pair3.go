package main

import "fmt"

type Reader interface {
	ReaderBook()
}

type Writer interface {
	WriterBook()
}

//具体类型
type Book struct {
}

func (this *Book) ReaderBook() {
	fmt.Println("Read a Book")
}

func (this *Book) WriterBook() {
	fmt.Println("write a Book")
}

func main()  {
	//b: pair<type: Book, value: Book{}地址>
	b := &Book{}

	//r: pair<type: , value: >
	var r Reader
	//r: pair<type: Book, value: Book{}地址>
	r = b

	r.ReaderBook()

	var w Writer
	//r: pair<type: Book, value: Book{}地址>
	w = r.(Writer)  //此处的断言为什么会成功，因为w r 具体的type是一致的

	w.WriterBook()
}