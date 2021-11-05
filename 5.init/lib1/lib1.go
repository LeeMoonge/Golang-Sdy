package lib1

import "fmt"

func Lib1Test() {
	fmt.Println("in LibTest func ...")
}

func init() {
	fmt.Println("in lib1 init func ...")
}