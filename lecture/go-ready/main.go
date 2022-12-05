package main

import (
	"fmt"
)

// func main() {
// 	// var name string
// 	// fmt.Scan(&name)
// 	// fmt.Print(name)

// 	// var name string
// 	// var age int

// 	// fmt.Println("이름을 입력하세요.")
// 	// fmt.Print("이름 입력: ")
// 	// fmt.Scanln(&name)

// 	// fmt.Println("나이를 입력하세요.")
// 	// fmt.Print("나이 입력: ")
// 	// fmt.Scan(&age, &name)

// 	// fmt.Println("이름: ", name, "\n", "나이: ", age)

// 	// 이름을 담을 변수 선언
// 	var name string

// 	// input을 name에 담아줌
// 	fmt.Scan(&name)

// 	// Sprintln으로 문자열 핸들링
// 	hello := fmt.Sprintln("Hi, my name is", name)

// 	// print로 hello 문자열 출력
// 	fmt.Print(hello)
// }

// func main() {
// 	var name, addr string
// 	var num int
// 	fmt.Println("이름 번호 주소")
// 	var re int
// 	re, _ = fmt.Scanln(&name, &num, &addr)
// 	fmt.Println("이름은 ", name, " 번호는 ", num)
// 	fmt.Println("주소는 ", addr)
// 	fmt.Println(re)
// }

// func multiCalc(a, b int)(x, y, z int) {
// 	return (a + b), (a * b), (a / b)
// }

// func main() {
// 	i, j, k := multiCalc(20, 10)
// 	fmt.Printf("Sum :%d, Mul : %d, Div : %d \n", i, j, k)
// }

// func multiCalc(a, b int)(x, y, z int) {
// 	return (a + b), (a * b), (a / b)
// }

// func main() {

// 	// i, j, k := multiCalc(20, 10)
// 	// fmt.Printf("Sum :%d, Mul : %d, Div : %d \n", i, j, k)

// 	// _, j, k := multiCalc(20, 10)
// 	// fmt.Printf("Mul : %d, Div : %d \n", j, k)

// 	// i, _, k := multiCalc(20, 10)
// 	// fmt.Printf("Sum :%d, Div : %d \n", i, k)

// 	// i, _, _ := multiCalc(20, 10)
// 	// fmt.Printf("Sum :%d \n", i)
// }

// func one() {
// 	fmt.Println("one")
// }

// func two() {
// 	fmt.Println("two")
// }

// func main() {
// 	// 파일을 생성, 이 때 첫번째 return은 파일을 가리키는 포인터이다.
// 	file, _ := os.Create("./res1.txt")
// 	// defer를 사용해 파일을 닫는 작업을 모든 작업 이후 처리하게 해줌
// 	defer file.Close()

// 	// buffer 작업 최적화를 위해 w라는 변수에 file 주소값을 담아줌
// 	w := bufio.NewWriter(file)
// 	// 버퍼에서 w를 지워줌
// 	defer w.Flush()

// 	// 실제 w(file)에 뒤에 오는 string을 입력해줌
// 	fmt.Fprint(w, "Fprint sprint")
// }

// 가변함수 -> ...부분은 배열로 취급이 되는구나
// func tempFunc(i ...string) {
// 	fmt.Println(i) // [hi hello 방가방가]
// 	fmt.Println(i[1]) // hello
// }

// func main() {
// 	tempFunc("hi", "hello", "방가방가")
// }

// func test(x int, y string) (a int, b string) {
// 	return x, y
// }

// func main() {
// 	// i, j := test(1, "35")
// 	// fmt.Printf("%d, , , , , %s \n" , i, j)

// 	// fmt.Println(reflect.TypeOf(i))

// 	// k, s := "hi", "hello"
// 	// fmt.Println(k)
// 	// fmt.Println(s)

// 	var i int
// 	fmt.Println(&i)
// }

// func main() {
// 	// var a1 []int
// 	// var a2 [5]int

// 	// fmt.Println(reflect.ValueOf(a1).Kind())
// 	// fmt.Println(reflect.TypeOf(a1))
// 	// fmt.Println(reflect.ValueOf(a2).Kind())
// 	// fmt.Println(reflect.TypeOf(a2))

// 	var intSlice = []int{0, 1, 2, 3}
// 	fmt.Printf("%#v\n", intSlice)
// 	fmt.Printf("%v\n", intSlice)
// 	fmt.Printf(":2 %v\n", intSlice[:2])
// 	fmt.Printf("1:3 %v\n", intSlice[1:3])
// 	fmt.Printf("2: %v\n", intSlice[2:])
// 	fmt.Printf("0:3 %v\n", intSlice[0:3])

// }

// 메모리 할당 및 초기화

// func NewFile(fd int, name string) *File {
// 	if fd < 0 {
// 		return nil
// 	}
// 	f := new(File)
// 	f.fd = fd
// }

// func main() {
// 	var strArray1 [5] string
// 	strArray1[0] = "a"
// 	strArray1[1] = "b"
// 	strArray1[2] = "c"
// 	strArray1[3] = "d"
// 	strArray1[4] = "e"
// 	fmt.Print(strArray1)

// 	var strArray2 [5]string = [5]string {"1", "2", "3", "4", "5"}
// 	fmt.Println(strArray2)

// 	strArray3 := []string{"10", "20", "30"}
// 	fmt.Println(strArray3)

// 	strArray4 := [...]string{"100", "200", "300"}
// 	fmt.Println(strArray4)
// 	fmt.Println(len(strArray4))

// 	strArray5 := [10]string{2: "codz", 4: "states"}
// 	fmt.Println(strArray5)
// 	fmt.Printf("%v\n", strArray5)
// }

// type File struct {
// 	fd int
// 	name string
// 	dirinfo string
// 	nepipe int
// }

// func main() {
// 	Fpointer, _ := os.Create("./res2.txt")
// 	defer Fpointer.Close()

// 	fmt.Fprintln(Fpointer, "text text");

// 	Fpointer.

// 	fmt.Println("hi")
// }

// func main() {
// 	// var map1 map[string]int

// 	// map2 := make(map[string]int)
// 	// map3 := make(map[string]string, 2)

// 	var timeZone = map[string]int {
// 		"UTC" : 0,
// 		"EST" : 5,
// 		"CST" : 6,
// 		"MST" : 7,
// 		"PST" : 8,
// 	}

// 	timeSet := timeZone["CST"]

// 	fmt.Println(timeSet)

// 	for key, value := range timeZone {
// 		fmt.Println(key, value)
// 	}

// temp := make(map[string]int)

// temp["hi"] = 1
// temp["hello"] = 2
// temp["bye"] = 3

// for key, value := range temp {
// 	fmt.Println("key: ", key, "value: ", value)
// }

// }

// func main() {
// 	f, _ := os.Open("res.txt")

// 	fmt.Println(f)

// 	var intSlice = []int{0, 1, 2, 3}
// 	fmt.Printf("%#v\n", intSlice)
// 	fmt.Printf("%v\n", intSlice)
// 	fmt.Printf(":2 %v\n", intSlice[:2])
// 	fmt.Printf("1:3 %v\n", intSlice[1:3])
// 	fmt.Printf("2: %v\n", intSlice[2:])
// 	fmt.Printf("0:3 %v\n", intSlice[0:3])

// }

// func main() {
// 	for i := 0; i < 100; i++ {
// 		fmt.Printf("%d", i)
// 	}

// 	sum := 0
// 	i := 100

// 	nums := [4]int{99, 98, 98, 32}
// 	for el, el2 := range nums {
// 		fmt.Println(el, el2)
// 	}

// 	var arr [10]int

// 	// 배열이기때문에 return의 첫번째 값이 index이다.
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
// 	x := getValue(10) // 출력: x == 10
// 	y := 20 // 출력: y == 20

// 	switch x {
// 	case 10:
// 		fmt.Println("x == 10")
// 	// golang은 컴파일러가 case문마다 자동으로 break문을 추가해준다.
// 	case 20:
// 		fmt.Println("x == 20")
// 	}

// 	switch y {
// 	case 10:
// 		fmt.Println("y == 10")
// 	// golang은 컴파일러가 case문마다 자동으로 break문을 추가해준다.
// 	case 20:
// 		fmt.Println("y == 20")
// 	}
// }

// func main() {
// 	// if문 내에서 사용할 변수를 생성하여 적용해 줄 수 있다.
// 	if x := getValue(10); x == 10 {
// 		fmt.Println("x == 10")
// 	} else {
// 		fmt.Println("x ==", x)
// 	} // 출력: x == 10

// 	if x := getValue(20); x == 10 {
// 		fmt.Println("x == 10")
// 	} else {
// 		fmt.Println("x ==", x)
// 	} // 출력: x == 20

// 	// 위 if문 내에서 사용한 x는 if - else 스코프를 벗어나는 순간 사라지므로
// 	// 아래와 같이 전역적으로 사용할 수 없다.
// 	fmt.Println(x) // 컴파일 에러
// }

// func main() {
// 	r := gin.Default()
// 	r.GET("/", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"result" : "ok",
// 		})
// 	})
// 	r.Run()
// }

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

// func main() {
// 	f, err := os.Open("es2.txt")
// 	defer f.Close()

// 	if err != nil {
// 		panic(errors.New("file not exist"))
// 	}

// 	result, error := ioutil.ReadFile("res2.txt")
// 	check(error)
// 	fmt.Println(string(result))
// }

// func calc(a int) {
// 	defer func ()  {
// 		if r := recover(); r != nil {
// 			fmt.Println(r)
// 		}
// 	}()
// 	if a == 0 {
// 		panic(errors.New("a is zero"))
// 	} else {
// 		fmt.Println(a)
// 	}
// }

// func main() {
// 	calc(1)
// 	calc(2)
// 	calc(0)
// 	calc(3)
// 	calc(4)
// }

// func counter(s string, n int) {
// 	for i := 0; i < n; i++ {
// 		time.Sleep(time.Second)
// 		fmt.Println(s, " : ", i)
// 	}
// }

// func main() {
// 	fmt.Println("start go")

// 	go func() {
// 		defer counter("e", 10)
// 	}()

// 	go counter("a", 10)

// 	go func() {
// 		counter("b", 10)
// 	}()

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			time.Sleep(time.Second)
// 			fmt.Println("c : ", i)
// 		}
// 	}()

// 	// var i int = 0
// 	// for {
// 	// 	time.Sleep(time.Second)
// 	// 	fmt.Println("c : ", i)
// 	// 	i++
// 	// }

// 	counter("d", 10)
// }




func main() {
	fmt.Println("Hello World")
}