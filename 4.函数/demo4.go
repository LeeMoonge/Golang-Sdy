package main

import "fmt"

//func 函数名(参数1 类型， 参数二 类型) 返回值类型 {
//}
func f1(a string, b int) int {
	fmt.Println("in func f1")
	fmt.Printf("a = %v, b = %v\n", a, b)

	c := 777
	return c
}

//返回多个返回值，返回值匿名
func f2(c int, d string) (int, string) {
	fmt.Println("in func f2")
	fmt.Println("c =", c, "d =", d)
	re1, re2 := 666, "str666"
	return re1, re2
}

//返回多个返回值，有形参名称
func f3(e int, f string) (r1 int, r2 string) {
	fmt.Println("in func f3")
	fmt.Println("e =", e, "f =", f)
	r1 = e + 1
	r2 = f
	return
}

//返回多个返回值，且返回值的数据类型一样
func f4(g int, h string) (r1, r2 int) {
	fmt.Println("in func f4")
	fmt.Println("g =", g, "h =", h)
	fmt.Println("r1 =", r1, "r2 =", r2)

	return 100, 200
}

func main()  {
	c := f1("aaa", 666)
	fmt.Println("c =", c)

	f2_ret1, f2_ret2 := f2(999, "google")
	fmt.Println("f2_ret1 =", f2_ret1, "f2_ret2 =", f2_ret2)

	r1, r2 := f3(100, "golang")
	fmt.Println("r1 =", r1, "r2 =", r2)

	f4_r1, f4_r2 := f4(100, "golang_test")
	fmt.Println("f4_r1 =", f4_r1, "f4_r2 =", f4_r2)
}
