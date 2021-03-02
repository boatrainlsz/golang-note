package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	//数组的range
	//_表示空白符，表示忽略下标
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for index, num := range nums {
		fmt.Println("index:", index, "num:", num)
	}
	//map的range
	kvMap := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvMap {
		println("k:", k, "v:", v)
	}

	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
