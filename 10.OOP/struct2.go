package main

import "fmt"

//如果累名首字母大写，表示其他包也能够访问
type Hero struct {
	//如果说类的属性首字母大写，表示该属性是对外能够访问的，否则的话只能怪类的内部访问
	Name string
	Age int
	Level int
}

func (this *Hero)Showhero()  {
	fmt.Println("hero name =", this.Name)
	fmt.Println("hero age =", this.Age)
	fmt.Println("hero level =", this.Level)
}

func (this *Hero) Getname()  {
	fmt.Println("hero name =", this.Name)
}

func (this *Hero)Setname(name string)  {
	//this 是调用该方法的对象的一个副本（拷贝）
	//建议使用将该对象的指针
	this.Name = name
}

func main()  {
	//创建一个对象
	hero := Hero{Name: "yqq", Age: 22, Level: 1}
	hero.Getname()
	hero.Showhero()

	name := "小紫"
	hero.Setname(name)

	hero.Getname()
	hero.Showhero()
}