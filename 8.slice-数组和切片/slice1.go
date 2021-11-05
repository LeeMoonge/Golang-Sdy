package main

import "fmt"

func checkSlice(myslice []int) bool {
	if myslice == nil{
		fmt.Println("这是一个空切片")
		return true
	} else {
		fmt.Println("该slice有空间")
		return false
	}
}


func main()  {
	//声明slice1是一个切片，并且初始化，初始化值是 1，2，3。长度len是3
	slice1 := []int{1, 2, 3}
	fmt.Printf("slice1 len=%d, array=%v, type=%T\n", len(slice1), slice1, slice1)


	//声明myArray2是一个切片，但是没有给myArray2分配空间
	var myArray2 []int
	fmt.Printf("myArray len=%d, array=%v", len(myArray2), myArray2)
	//没有默认值的不定长数组的默认值长度为0，需要重新给该变量开辟空间，使用make([]int, 长度)来重新开辟空间
	myArray2 = make([]int, 3)  //如果不重新开辟空间运行会报错
	myArray2[0] = 1
	fmt.Printf("myArray len=%d, array=%v, type=%T\n", len(myArray2), myArray2, myArray2)

	//声明slice2是一个切片，同时给myArray2分配空间，3个空间，初始化值是0
	var slice2 []int = make([]int, 3)
	fmt.Printf("slice2 len=%d, array=%v, type=%T\n", len(slice2), slice2, slice2)

	//声明slice3是一个切片，同时给myArray2分配空间，3个空间，初始化值是0，通过 := 推导出myArray是一个切片
	slice3 := make([]int, 4)
	fmt.Printf("slice3 len=%d, array=%v, type=%T\n", len(slice3), slice3, slice3)

	//判断一个slice是否为0
	result := checkSlice(slice1)
	fmt.Println(result)

}
