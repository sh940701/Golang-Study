package main

import (
	"fmt"
)

// func main() {
// 	var name string
// 	var age int

// 	fmt.Println("이름을 입력하세요 : ")
// 	fmt.Print("이름 입력")
// 	fmt.Scan(&name)

// 	fmt.Println("나이를 입력하세요 : ")
// 	fmt.Print("나이 입력")
// 	fmt.Scan(&age)

// 	fmt.Printf("이름: %s, 나이: %d세\n", name, age)
// }

// func main() {
// 	var name string
// 	var address string
// 	var age int

// 	// 공백, 개행을 기준으로 다음 값임을 인지한다.
// 	fmt.Scan(&name, &address, &age)

// 	// 포맷 지정자를 이용하여 개발자가 원하는 형태로 scan한다.
// 	fmt.Scanf("%s \n %s %d", &name, &address, &age)

// 	// 공백만을 기준으로 다음 값임을 인지한다.
// 	fmt.Scanln(&name, &address, &age)

// 	fmt.Println(name, address, age)

// 	// 김성현 서울시 27 -> 위 코드를 기준으로 공백만으로 구분하여 입력했을 경우 출력값
// 	// fmt.Scan -> 김성현 서울시 27
// 	// fmt.Scanf -> 김성현  0
// 	// fmt.Scanln -> 김성현 서울시 27

// 	// 김성현
// 	// 서울시
// 	// 27
// 	// 개행으로 구분하여 입력했을 경우 출력값

// 	// fmt.Scan -> 김성현 서울시 27
// 	// fmt.Scanf -> 김성현 서울시 0
// 	// fmt.Scanln -> 김성현 0
// }

// func main() {
// 	var i int
// 	var s string

// 	i10 := 10
// 	s5 := "abcde"

// 	fmt.Println(reflect.TypeOf(i)) // int
// 	fmt.Println(reflect.TypeOf(s)) // string

// 	fmt.Println(reflect.TypeOf(i10)) // int
// 	fmt.Println(reflect.TypeOf(s5)) // string
// }

// func resInt(n int) int {
// 	return n
// }

// func resStr(s string) string {
// 	return s
// }

// func main() {
// 	// var i int
// 	// var j int = 10
// 	// i = 10
// 	// var s string = "name"
// 	// var h = 20
// 	// var str = "codz"
// 	// k := 13
// 	// b := true
// 	// b = false
// 	// const ss string = "name"
// 	// const ii = 15
// 	// const s1 = "name"
// 	// const i1 int = 25

// 	l := resInt(30)
// 	m := resStr("codz")

// 	fmt.Println(l) // 30
// 	fmt.Println(m) // codz

// 	fmt.Println(reflect.TypeOf(l)) // int
// 	fmt.Println(reflect.TypeOf(m)) // string
// }

//문제1 두 인자의 합과 곱을 리턴하는 multiCalc 함수 부분을 채우세요
// func multiCalc(a, b int)(x, y, z int) {
// 	return a + b, a * b, a / b
// }

// func main() {
// 	add, mul, div := multiCalc(20, 10)
// 	fmt.Printf("Add :%d, Mul : %d, Div : %d", add, mul, div) // Sum :30, Mul : 200, Div : 2
// }

// func main() {
// 	_, a, b := multiCalc(20, 10)
// 	c, _, d := multiCalc(20, 10)
// 	e, f, _ := multiCalc(20, 10)
// 	g, _, _ := multiCalc(20, 10)

// 	fmt.Println(a, b, c, d, e, f, g) // 200, 2, 30, 2, 30, 200, 30
// }

// func one() {
// 	fmt.Println("1111")
// }

// func two() {
// 	fmt.Println("2222")
// }

// func main() {
//   defer two()
//   one()
// }

// func main() {
// 	// 새로운 읽기 전용 객체 생성
// 	input := bufio.NewReader(os.Stdin)
// 	// 개행문자를 구분자로 읽어오도록 설정
// 	line, _ := input.ReadString('\n')

// 	// ./경로에 res.txt라는 파일 생성
// 	file, err := os.Create("./res.txt")

// 	// 에러 발생시 처리
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	// 함수 종료 후 file에 대한 close함수 실행
// 	defer file.Close()

// 	// 파일에 대한 버퍼를 만들어줌(?)
// 	w := bufio.NewWriter(file)
// 	// 버퍼링된 데이터를 IO에 기록(?)
// 	defer w.Flush()

// 	// w에 str을 기록
// 	fmt.Fprintln(file, line)
// }


func temp(s ...string) {
	// 가변인자로 이루어진 s는 slice 이다.
	// s에 대한 range문은 인덱스와 값을 리턴한다.
	for _, value := range s {
		fmt.Println(value)
	}
	fmt.Println("")
}

func main() {
	temp("ab")
	temp("ab", "cd")
	temp("ab", "cd", "ef")
}
/*
	ab

	ab
	cd

	ab
	cd
	ef
*/