package main

import (
	"Golang-Sdy/5.init/lib1"
	"Golang-Sdy/5.init/lib2"
)

/*import (
	_ "Golang-Sdy/5.init/lib1"  //在包名前给包起一个别名，使用 _ 时不能使用改包中的所有方法，但是会在导包时执行该包的 init 初始化方法
	my_lib2 "Golang-Sdy/5.init/lib2"  //给包起一个别名，使用别名时可以正常使用该包中的所有方法，也会执行该包的 init 方法初始化
	. "Golang-Sdy/5.init/lib2"  //（不推荐）给包起一个为.的别名，可以直接使用该包中的所有方法，使用时直接用方法名即可，注意：如果有方法名同名容易冲突
)*/

func main()  {
	lib1.Lib1Test()
	lib2.Lib2Test()
}
