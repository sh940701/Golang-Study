package main

import (
	"errors"
	"fmt"
)

// func main() {
// 	i := 1
// 	for ; i <= 10; i++ {
// 		fmt.Println(i)
// 	}

// 	i = 1
// 	for i <= 10 {
// 		fmt.Println(i)
// 		i++
// 	}

// 	for i := 1; ; i++ {
// 		fmt.Println(i)
// 		if i == 10 {
// 			break
// 		}
// 	}

// 	for i := 0; i < 100; i++ {
// 		fmt.Printf("%d", i)
// 	}
// }

// func main() {
// 	var arr [10]int

// 	for idx0, _ := range arr {
// 		arr[idx0] = len(arr) - idx0
// 	}

// 	for idx1, value1 := range arr {
// 		fmt.Println("idx: ", idx1, "value: ", value1)
// 	}
// }

// func getValue(value int) (int) {
// 	return value
// }

// func main() {
// 	if x := getValue(10); x == 10 {
// 		fmt.Println("x == 10")
// 	} else {
// 		fmt.Println("x == ", x)
// 	}

// 	if x := getValue(30); x == 10 {
// 		fmt.Println("x == 10")
// 	} else if x == 20 {
// 		fmt.Println("x == 20")
// 	} else {
// 		fmt.Println("x == ", x)
//   	}
// }

// func getValue(value int) (int) {
// 	return value
// }

// func main() {
// 	x := getValue(10)
// 	y := 20

// 	switch x {
// 	case 10:
// 		fmt.Println("x == 10")
// 	case 20:
// 		fmt.Println("x == 20")
// 	}

// 	switch y {
// 	case 10:
// 		fmt.Println("y == 10")
// 	case 20:
// 		fmt.Println("y == 20")
// 	}
// }

// func main() {
// 	var i int
// 	for i = 0; i < 10; i++{
// 		fmt.Printf("%d", i)
// 	}

// 	fmt.Println("")
// 	if i == 10 {
// 		fmt.Println("if - else : i == 10")
// 	} else {
// 		fmt.Println("if - else : i != 10")
// 	}

// 	switch i {
// 	case 10:
// 		fmt.Println("switch : i == 10")

// 	case 9:
// 		fmt.Println("switch : i != 10")
// 	}

// }

// func main() {
// 	f, err := os.Open("/.string.txt")
// 	defer f.Close()

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Fprint(f, "hello")
// }


var res = 1

func Calc(n int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic 복구 - ", r)
		}
	}()
	if n == 0 {
		panic(errors.New("zero calc"))
	} else {
		res *= n
		fmt.Println("Result: ", res)
	}
}

func main() {
	Calc(1)
	Calc(2)
	Calc(0)
	Calc(3)
}