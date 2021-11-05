package main

import (
	"fmt"
)

//参数类型指明传入数组类型长度，就必须传入相同类型或长度的固定长度数组，不然会报错
func printArray(myArray [4]int)  {
	for index, value := range myArray{
		fmt.Printf("===in printArray func index=%v, value=%v===\n", index, value)
	}
	//固定长度数组传参为值传递，是将值赋值给函数调用时的参数，在函数体内部改变了数组的值，外面的数组不会发生改变
	myArray[0] = 999
}


func main()  {
	//固定长度的数组
	//类型一
	var myArray1 [10]int
	for i := 0;i < len(myArray1);i++{
		fmt.Printf("myArray1 index:i=%v value=%v\n", i, myArray1[i])
	}

	fmt.Println("================================")

	//类型二
	myArray2 := [10]int{1, 2, 3, 4}
	for index, valve := range myArray2{
		fmt.Printf("myArray2 index=%v value=%v\n", index, valve)
	}

	fmt.Println("================================")

	//类型三
	myArray3 := [4]int{11, 22, 33, 44}
	for index, value := range myArray3{
		fmt.Printf("myArray3 index=%v, value=%v\n", index, value)
	}


	//查看数组的数据类型
	//fmt.Printf("myArray1 type : %T\n", myArray1)
	//fmt.Printf("myArray2 type : %T\n", myArray2)
	fmt.Printf("myArray3 type : %T\n", myArray3)

	printArray(myArray3)

	for index, value := range myArray3{
		fmt.Printf("myArray3 index=%v, value=%v\n", index, value)
	}
}
