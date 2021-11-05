package main

import "fmt"

/*
	切片的取值
*/

func main()  {
	s := []int{1, 2, 3, 4}  // s = [1, 2, 3, 4]

	s1 := s[0:2]
	fmt.Printf("s1 = %v, type:%T\n", s1, s1)

	s2 := s[0:3]
	fmt.Printf("s2 = %v, type:%T\n", s2, s2)

	//s1[0] = 100
	//fmt.Printf("s1 = %v, type:%T\n", s1, s1)
	//fmt.Printf("s2 = %v, type:%T\n", s2, s2)

	//copy可以将底层的slice一起进行拷贝
	s3 := make([]int, 3)
	//将s中的值依次拷贝到s3中，因为创建s3时定义的长度为3，所以只拷贝了s中的前三个元素
	copy(s3, s)
	fmt.Printf("s3 = %v, type:%T\n", s3, s3)
}
