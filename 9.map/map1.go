package main

import "fmt"

func main()  {
	//====第一种声明方式
	//声明myMap是一种map类型，key是string，value是string
	var myMap1 map[string]string
	if myMap1 == nil{
		fmt.Println("myMap 为空map")
	}
	//在使用前，需要用make()给新创建的map分配数据空间
	myMap1 = make(map[string]string, 10)
	myMap1["SHENZHEN"] = "one"
	myMap1["BEIJING"] = "two"
	myMap1["SHANGHAI"] = "three"
	fmt.Println("myMap1 = ", myMap1)

	//====第二种声明方式
	myMap2 := make(map[int]string)
	myMap2[1] = "aab"
	myMap2[2] = "bba"
	myMap2[3] = "aabb"
	fmt.Println("myMap = ", myMap2)

	//====第三种声明方式
	myMap3 := map[string]string{
		"name": "xxbb",
		"age": "10",
		"birthday": "2021-1-1",  //map类型的最后一行也需要加逗号
	}
	fmt.Println("myMap3 = ", myMap3)
}
