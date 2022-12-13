package Students

import (
	"fmt"
)

type Students struct {
}

func (s Students) Add() {
	fmt.Println("Add Student")
}
func (s Students) Delete() {
	fmt.Println("Delete Student")
}