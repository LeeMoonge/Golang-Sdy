package main

import (
	"fmt"
	"reflect"
)

//结构体标签

type resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex string `info:"sex"`
}

func getTag(obj interface{})  {
	a := reflect.TypeOf(obj).Elem()  //.Elem() 当前结构体全部的元素

	for i := 0; i < a.NumField(); i ++ {
		tagInfo := a.Field(i).Tag.Get("info")
		tagDoc := a.Field(i).Tag.Get("doc")
		fmt.Println("TagInfo = ", tagInfo, "tagDoc = ", tagDoc)
	}
}

func main()  {
	test := resume{Name: "apdl", Sex: "男"}
	getTag(&test)
}
