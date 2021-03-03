package main

import "fmt"

//用法3，全局匿名函数
var (
	func1 = func(n1 int, n2 int) int { return n1 * n2 }
)

func main() {
	//用法1,只调用一次
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 30)
	fmt.Println("res1=", res1)

	//用法2,赋给变量a，多次diaoyong
	a := func(n1 int, n2 int) int {
		return n1 + n2
	}
	res2 := a(10, 2)
	res3 := a(10, 19)
	fmt.Println("res2=", res2)
	fmt.Println("res3=", res3)

	//全局匿名函数
	res4 := func1(1, 5)
	fmt.Println("res4=", res4)

}
