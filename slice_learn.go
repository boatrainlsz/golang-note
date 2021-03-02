package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(len(numbers), cap(numbers))
	fmt.Println(numbers[:])
	fmt.Println(numbers[1:4])
	fmt.Println(numbers[:4])
	fmt.Println(numbers[4:])
	numbers1 := numbers[:3]
	fmt.Println(numbers1)

	//append
	var numbers2 []int
	numbers2 = append(numbers2, 0)
	numbers2 = append(numbers2, 1, 2, 23)
	copy(numbers1, numbers2)
}
