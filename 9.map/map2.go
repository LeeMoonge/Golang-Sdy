package main

import "fmt"


func printMap(myMap map[string]string)  {
	for key, value := range myMap{
		fmt.Println("key =", key, "value =", value)
	}
}

func updateMap(myMap map[string]string)  {
	//myMap是一个引用传递，默认将传入参数的指针传进来
	myMap["UK"] = "London"
}

func main()  {
	myMap1 := make(map[string]string)

	//添加新的键值对
	myMap1["China"] = "beijing"
	myMap1["Japan"] = "tokyo"
	myMap1["USA"] = "NewYork"

	//遍历myMap1, 打印其键值对
	printMap(myMap1)

	fmt.Println("======================")

	//删除操作
	delete(myMap1, "China")

	//修改操作
	myMap1["USA"] = "DC"

	updateMap(myMap1)

	//遍历myMap1, 打印其键值对
	printMap(myMap1)
}
