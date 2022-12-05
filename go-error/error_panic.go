package main

import "fmt"

// panic
// func divide(a, b int) {
// 	if b == 0 {
// 		panic("B cannot be a zero")
// 	}
// 	fmt.Println(a / b)
// }

// func main() {
// 	divide(9, 3)
// 	divide(3, 0)
// }

// recover
func f() {
	fmt.Println("f()함수 시작")

	// defer함수는 함수가 종료될 때 실행된다.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic 복구 -", r)
		}
	}()
	
	// g() 함수를 실행한다.
	// 이 때 g()함수 하위의 h()함수에서 panic이 났고, f()함수에서 recover이 이루어졌다.
	// panic의 발생으로 역으로 함수를 파괴하면서 올라오고 있었기 때문에 g()함수를 호출한 상태로
	// f()함수가 종료되어 아래 "f()함수 끝"은 출력되지 않는다. 다만 함수가 아예 종료되는 시점에
	// stack에 있던 defer함수는 실행되기 때문에 출력은 안되지만 recover은 실행된것이다.
	g()
	fmt.Println("f()함수 끝")
}

func g() {
	// g() 함수에서는 h 함수를 실행한다.
	fmt.Printf("9 / 3 = %d \n", h(9, 3))
	fmt.Printf("9 / 3 = %d\n", h(9, 0))
}

func h(a, b int) int {
	// h 함수는 내부적으로 b가 0이면 panic을 실행한다.
	if b == 0 {
		panic("b cannot be zero")
	}
	return a / b
}

func main() {
	// main에서 f() 함수를 실행한다.
	f()
	fmt.Println("프로그램이 계속 실행됨")
}

/*
	f()함수 시작 -> f() -> g() -> h() 순서로 실행된다.

	9 / 3 = 3 -> h(9, 3) 이 실행되었다.

	panic 복구 - b cannot be zero -> h(9, 0)에 대해 panic이 실행되었다.

	프로그램이 계속 실행됨 -> f()함수까지는 파괴되었지만, 
	recover로 인해서 main함수는 이후 로직을	계속 이어가게 된다.
*/