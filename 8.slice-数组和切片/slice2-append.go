package main

import "fmt"

/*
	切片的追加
*/

func main()  {
	var numbers []int = make([]int, 3, 5)
	fmt.Printf("numbers len=%d, cap=%d, slice=%v, type=%T\n", len(numbers), cap(numbers), numbers, numbers)

	//向numbers切片追加一个元素1，numbers：len=4，cap=5，slice=[0, 0, 0, 1], type=[]int
	numbers = append(numbers, 1)
	fmt.Printf("numbers len=%d, cap=%d, slice=%v, type=%T\n", len(numbers), cap(numbers), numbers, numbers)

	//向numbers切片追加一个元素2，numbers：len=5，cap=5，slice=[0, 0, 0, 1, 2], type=[]int
	numbers = append(numbers, 2)
	fmt.Printf("numbers len=%d, cap=%d, slice=%v, type=%T\n", len(numbers), cap(numbers), numbers, numbers)

	//向一个容量已经满了的的slice追加元素: 不会报错，而是将该slice的初始cap(容量)*2，继续新增元素
	numbers = append(numbers, 3)
	//numbers len=6, cap=10, slice=[0 0 0 1 2 3], type=[]int（初始的cap(容量)为5，此时新增超过容量的元素，初始容量*2，长度依次增加）
	fmt.Printf("numbers len=%d, cap=%d, slice=%v, type=%T\n", len(numbers), cap(numbers), numbers, numbers)

	fmt.Println("======================================================")

	var numbers2 []int = make([]int, 3)
	fmt.Printf("numbers2 len=%d, cap=%d, slice=%v, type=%T\n", len(numbers2), cap(numbers2), numbers2, numbers2)
	numbers2 = append(numbers2, 1)
	fmt.Printf("numbers2 len=%d, cap=%d, slice=%v, type=%T\n", len(numbers2), cap(numbers2), numbers2, numbers2)
}


