package main

import "fmt"

func main() {
	var a int = 20

	var ip = &a

	fmt.Println(ip)
	fmt.Println(*ip)
}
