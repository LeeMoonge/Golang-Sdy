package main

import "fmt"

//interface本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string  //重写该方法时如果函数有返回值，必须要声明类型，不然会报错
	GetType() string
}

//具体的类
type Cat struct {
	Color string
}

func (this *Cat) Sleep()  {
	fmt.Println("Cat is Sleep ...")
}

func (this *Cat) GetColor() string {
	return this.Color
}

func (this *Cat) GetType() string {
	return "Cat"
}

//具体的类
type Dog struct {
	Color string
}

func (this *Dog) Sleep()  {
	fmt.Println("Dog is Sleep ...")
}

func (this *Dog) GetColor() string {
	return this.Color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func ShowAnimal(animal AnimalIF){
	animal.Sleep()
	fmt.Println("animal Color =", animal.GetColor())
	fmt.Println("animal Type =", animal.GetType())
}

func main()  {
	/*var animal AnimalIF
	animal = &Cat{"Green"}
	animal.Sleep()
	fmt.Println("animal Color:", animal.GetColor())
	fmt.Println("animal Type:", animal.GetType())

	animal = &Dog{"Black"}
	animal.Sleep()
	fmt.Println("animal Color:", animal.GetColor())
	fmt.Println("animal Type:", animal.GetType())
	*/

	cat := Cat{"white"}
	dog := Dog{"black"}

	ShowAnimal(&cat)
	ShowAnimal(&dog)
}