package main

import (
	"errors"
	"fmt"
	"math"
)

/**
https://www.runoob.com/go/go-error-handling.html
*/

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}
func main() {
	_, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	}
}
