package main

import (
	"fmt"
	"reflect"
)

// func main() {
// 	var strArray1 [5]string
// 	strArray1[0] = "a"
// 	strArray1[1] = "b"
// 	strArray1[2] = "c"
// 	strArray1[3] = "d"
// 	strArray1[4] = "e"

// 	fmt.Println(strArray1)
// 	//문제1 : 0~4까지 모두 채운후, 전체 코드와 해당 출력 결과는?

// 	var strArray2 [5]string = [5]string {"1", "2", "3", "4", "5"}
// 	fmt.Println(strArray2)
// 	//문제2 : xxxx 자리에 올바른 배열을 선언하고 전체 코드와 해당 출력 결과는?

// 	strArray3 := [5]string{"10", "20", "30"}
// 	fmt.Println(strArray3)
// 	//문제3 : 전체 코드와 해당 출력 결과는?

// 	strArray4 := [...]string{"100", "200", "300"}
// 	fmt.Println(strArray4)
// 	fmt.Println(len(strArray4))
// 	//문제4 : 전체 코드와 두 출력 결과는?

// 	strArray5 := [5]string{2:"codz", 4:"states"}
// 	fmt.Println(strArray5)
// 	fmt.Printf("%#v\n", strArray5)
// 	//문제5 : 전체 코드와 해당 출력 결과는?
// }

// func main() {
// 	// var a1 []int
// 	// var a2 [5]int

// 	// fmt.Println(reflect.ValueOf(a1).Kind())
// 	// fmt.Println(reflect.TypeOf(a1))

// 	// fmt.Println(reflect.ValueOf(a2).Kind())
// 	// fmt.Println(reflect.TypeOf(a2))
// 	//a : 출력의 결과값은?

// 	var intSlice = []int{0, 1, 2, 3}
// 	fmt.Printf("%#v\n", intSlice)
// 	fmt.Printf("%v\n", intSlice)
// 	fmt.Printf(":2 %v\n", intSlice[:2])
// 	fmt.Printf("1:3 %v\n", intSlice[1:3])
// 	fmt.Printf("2: %v\n", intSlice[2:])
// 	fmt.Printf("0:3 %v\n", intSlice[0:3])
// 	//b : 각 출력 결과값은?
// }

// func main() {
// 	// var map1 map[string]int

// 	// map2 := make(map[string]string)
// 	// map3 := make(map[int]string, 1000)

// 	// var timeZone = map[string]int{
// 	// 	"UTC": 0,
// 	// 	"EST": 5,
// 	// 	"CST": 6,
// 	// 	"MST": 7,
// 	// 	"PST": 8,
// 	// }

// 	// for key, value := range timeZone {
// 	// 	fmt.Printf("key: %s, value: %d", key, value)
// 	// 	fmt.Println("")
// 	// }

// 	// timeSet := timeZone["CST"]

// 	map0 := make(map[string]int)

// 	map0["a"] = 10
// 	map0["b"] = 20
// 	map0["c"] = 30

// 	for key, value := range map0 {
// 		fmt.Printf("%s - %d", key, value)
// 		fmt.Println("")
// 	}
// }

// type Persons struct {
// 	Name string
// 	Age int
// 	Pnum string
// }

// func main() {
// 	// var pers Persons
// 	// pers.Name = "codz"
// 	// pers.Age = 23
// 	// pers.Pnum = "01098765432"

// 	// fmt.Println(pers)
// 	// fmt.Printf("%#v\n", pers)
// 	// fmt.Println("name : ", pers.Name, " age : ", pers.Age, " pnum : ", pers.Pnum)

// 	p1 := new(Persons)
// 	p1.Name = "codz"
// 	p1.Age = 29
// 	p1.Pnum = "01098765432"
// 	fmt.Println("Person 1 : ", p1)
// 	fmt.Printf("%#v\n", p1)

// 	var p2 = new(Persons)
// 	p2.Name = "states"
// 	p2.Age = 33
// 	p2.Pnum = "01054329876"
// 	fmt.Println("Person 2 : ", p2)
// 	fmt.Printf("%#v\n", p2)

// 	//선언과 동시에 초기화
// 	var p3 = &Persons{"test", 17, "01045673210"}
// 	fmt.Println("Person 3 : ", p3)
// 	fmt.Printf("%#v\n", p3)

// 	//구조체 배열 형식
// 	p4 := []Persons{
// 		Persons{
// 			Name: "p1",
// 			Age:  10,
// 			Pnum: "9876",
// 		},
// 		Persons{
// 			Name: "p2",
// 			Age:  9,
// 			Pnum: "8765",
// 		},
// 	}
// 	fmt.Println(p4)
// 	fmt.Printf("%#v\n", p4)
// }

// func main() {
// 	var x interface{}
// 	x = 1
// 	fmt.Println(x)

// 	x = "Hello World"
// 	fmt.Println(x)
// }

// func main() {
// 	var x interface{}
// 	x = "10"

// 	sVal := x.(string)
// 	fmt.Println(sVal)
// 	x = 10
// 	nVal, bCnv := x.(int)
// 	fmt.Println(nVal, bCnv)
// }

// func main() {
// 	var x interface{}
// 	x = "hello"
// 	xType := reflect.TypeOf(x)

// 	if xType.Kind() == reflect.Int {
// 		fmt.Println("Int ", x)
// 	} else if xType.Kind() == reflect.String {
// 		fmt.Println("String ", x)
// 	}
// }

// type Persons struct {
// 	Name string
// 	Age int
// }

// func main() {
// 	codz := Persons{Name: "codz", Age: 20}

// 	var x interface{}
// 	x = codz

// 	xPerson := x.(Persons)
// 	// 여기서는 Age에 접근이 되는데, x 에 대해서는 직접적으로 Age 접근이 안됨
// 	// 그런데 reflect.TypeOf()를 해보면 x와 xPerson과 같은 결과가 나온다 -> main.Persons
// 	// 그럼에도 불구하고 타입 변환을 명시적으로 해줘야 멤버변수에 접근이 가능함...
// 	xPrsAge := x.(Persons).Age
// 	ty := reflect.TypeOf(xPerson)
// 	fmt.Printf("%v, %d, %s\n", xPerson, xPrsAge, xPerson.Name)
// 	fmt.Println(ty.String())
// }

// func main() {
// 	var personList = make(map[string]int)
// 	// element 2개 생성.
// 	personList["david"] = 20
// 	personList["json"] = 23

// 	// personList 출력
// 	fmt.Println(personList)

// 	// 2개 element 생성. 선언과 동시에 초기화
// 	var person = map[string] int {
// 		"david" : 20,
// 		"json" : 23,
// 	}

// 	// person 및 person의 길이 출력
// 	fmt.Println(person, len(person))

// 	// person element "c"의 value를 5 업데이트
// 	person["c"] = 5
// 	// person 출력
// 	fmt.Println(person)

// 	// person element 1개 삭제. 위에서 업데이트 안 한 element
// 	delete(person, "david")
// 	// person 출력
// 	fmt.Println(person)

// 	// 루프 - key,value 출력
// 	for key, value := range person {
// 		fmt.Println(key, value)
// 	}

// 	// person 및 personList 출력
// 	fmt.Println(person)
// 	fmt.Println(personList)
// }



type Person struct {
    Name string
    Age  int
    Pnum string
}

// make를 이용해 맵과 슬라이스 인스턴스 생성
func makeMap() map[string]int {
    return make(map[string]int)
}

func makeSlice() []int {
    return make([]int, 10)
}

func main() {
    p := new(Person)
    p.Age = 23

	fmt.Println(p) // &{ 23 }
	fmt.Println(reflect.TypeOf(p).Kind()) // ptr
	fmt.Println(reflect.TypeOf(p)) // *main.Person

    // print p => &{ 23 }
    // print ? => ptr
    // print ? => person

	
	m := makeMap()
	fmt.Println(m) // map[]
	fmt.Println(reflect.TypeOf(m).Kind()) // map
	fmt.Println(reflect.TypeOf(m)) // map[string]int

    // m := makeMap()
    // print m => map[]
    // print ? => map
    // print ? => map[string]int

	s := makeSlice()
	fmt.Println(s) // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(reflect.TypeOf(s).Kind()) // slice
	fmt.Println(reflect.TypeOf(s)) // []int


    // s := makeSlice()
    // print s => [ 0 0 0 0 0 0 0 0 0 0 ]
    // print ? => slice
    // print ? => []int

}