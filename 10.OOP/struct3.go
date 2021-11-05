package main

import "fmt"

//Human类
type Human struct {
	Name string
	Age int
}

func (this *Human)Eat()  {
	fmt.Println("Human Eat ...")
}

func (this *Human)Walk()  {
	fmt.Println("Human Walk ...")
}

//Super类
type SuperHuman struct {
	Human
	Level int
}

func (this *SuperHuman)Fly() {
	fmt.Println("SuperHuman Fly ...")
}

//重写父类Eat方法
func (this *SuperHuman)Eat() {
	fmt.Println("SuperHuman Eat ...")
}


//主函数
func main()  {
	var h SuperHuman
	h.Age = 20
	h.Name = "小青"
	h.Level = 1

	h.Fly()
	h.Walk()
	h.Eat()
}