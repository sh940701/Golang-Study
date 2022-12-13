package Person

import "fmt"

type Person struct {
}

func (p Person) Add() {
	fmt.Println("Add person")
}

func (p Person) Delete() {
	fmt.Println("Delete person")
}