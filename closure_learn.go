package main

import "fmt"

//funcGenerator()函数的返回值是函数，返回什么样的函数？返回一个入参和出参都是int的函数
func funcGenerator() func(int) int {
	var n int = 10
	//函数定义
	//这个匿名函数和匿名函数外的变量n形成一个整体，称为闭包
	return func(i int) int {
		n += i
		return n
	}
}
func main() {
	f := funcGenerator()
	fmt.Println(f(12))
	//注意闭包的定义，这里的输出不是23，而是35
	fmt.Println(f(13))
}
