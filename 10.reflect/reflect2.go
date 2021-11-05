package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func (u User) Call() {
	fmt.Println("User is called ...")
	fmt.Println(u)
}

func main()  {
	user := User{Id: 1, Name: "Abmzi", Age: 18}
	//user.Call()
	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{})  {
	//获取input的type
	inputType := reflect.TypeOf(input)
	fmt.Println("input type: ", inputType.Name())
	//获取input的value
	inputValue := reflect.ValueOf(input)
	fmt.Println("input value: ", inputValue)

	//通过type 获取里面的字段
	//1. 获取interface的reflect.Type, 通过Type得到NumField, 进行遍历
	//2. 得到每个field的数据类型
	//3. 通过field有一个Interface()方法得到对应的value
	for i := 0; i < inputType.NumField(); i ++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()

		fmt.Printf("field name=%s | field type=%v | field value=%v\n", field.Name, field.Type, value)
	}

	//通过type 获取里面的方法调用
	fmt.Println("inputType.NumMethod :", inputType.NumMethod())
	for i := 0; i < inputType.NumMethod(); i ++ {
		methodName := inputType.Method(i)
		fmt.Printf("method name=%v | method type=%v\n", methodName.Name, methodName.Type)
	}
}