package main

import (
	"fmt"
	ctl "lecture/go-debug/controller"
)

func main() {
	fmt.Println("Hello World")

	result := ctl.Sum(2, 3)

	fmt.Println(result)
}