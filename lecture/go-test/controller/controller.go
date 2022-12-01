package controller

import (
	"fmt"
)

func Sum(a, b int) int {
	fmt.Println("Sum: ", a, " + ", b)
	return a + b
}