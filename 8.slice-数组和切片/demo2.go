package main

import "fmt"

func printArray1(myArray []int)  {
	// _ 下划线表示匿名变量
	for _, value := range myArray{
		fmt.Printf("===in printArray1 func | value = %v===\n", value)
	}
	// 不定长数组的传参是引用传递，会将原数组的内存地址赋值给函数调用时的参数，在函数体中重新给数组内元素赋值，会改变原数组内元素的值
	myArray[0] = 100
}


func main()  {
	myArray1 := []int{1, 2, 3, 4} //动态数组，切片slice

	fmt.Printf("myArray1 = %v, type = %T\n", myArray1, myArray1)

	printArray1(myArray1)

	fmt.Printf("myArray1 = %v, type = %T\n", myArray1, myArray1)
	//print(myArray1 = [100 2 3 4], type = []int)

}