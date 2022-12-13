package main

// import "fmt"

// type Persons struct {
// }

// func (b Persons) Add() {
// 	fmt.Println("Add Person")
// }

// func (b Persons) Delete() {
// 	fmt.Println("Delete Person")
// }

// func (b *Persons) Alice() {
// 	b.name = "Alice"
// }

// func (b *Persons) Bob() {
// 	b.name = "Bob"
// }

// func main() {
// 	p := &Persons{"codz", 8, "0327894561"}
// 	fmt.Println(p)
// 	p.Alice()
// 	fmt.Println(p)
// 	p.Bob()
// 	fmt.Println(p)
// }

// type Persons struct {
// }

// func (b Persons) Add() {
// 	fmt.Println("Add Person")
// }

// func (b Persons) Delete() {
// 	fmt.Println("Delete Person")
// }

// type Persons struct {
// 	Name string
// 	Age int
// 	Pnum string
// }

// func (p Persons) Alice() { // Value 리시버
// 	p.Name = "alice"
// 	p.Age = 7
// 	p.Pnum = "01012345678"
// }

// func (p *Persons) Bob() { // Pointer 리시버
// 	p.Name = "bob"
// 	p.Age = 5
// 	p.Pnum = "01098765432"
// }

// func main() {
// 	p := &Persons{"codz", 8, "0327894561"}
// 	fmt.Println(p)
// 	p.Alice()
// 	fmt.Println(p)
// 	p.Bob()
// 	fmt.Println(p)
// }

// Persons "lecture/1213/controller/persons"
// import (
// 	"fmt"
// 	Person "lecture/1213/controller/person"
// 	Students "lecture/1213/controller/student"
// )

// type Persons struct {
// 	Students.Students
// 	Person.Person
// }

// func (p Persons) Add() {
// 	fmt.Println("Add Person")
// }

// func (p Persons) Delete() {
// 	fmt.Println("Delete Person")
// }

// func main() {
// 	var p Person.Person
// 	p.Add()
// 	p.Delete()

// 	var s Students.Students
// 	s.Add()
// 	s.Add()
// }

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name string
	Age int
	Pnum string
}

// func main() {
// 	p := Person{Name: "codz", Age: 23, Pnum: "01098765432"}
// 	jData, err := json.Marshal(p)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	sJData := string(jData)
// 	fmt.Println(sJData)
// }

// func main() {
// 	p := Person{Name: "codz", Age: 23, Pnum: "01098765432"}
// 	jData, _ := json.Marshal(p)
	
// 	var p2 Person
// 	json.Unmarshal(jData, &p2)

// 	fmt.Println(p2)
// 	fmt.Println(p2.Name)
// }


// func main() {
// 	p := Person{Name: "codz", Age: 23, Pnum: "01098765432"}
	
// 	encoder := json.NewEncoder(os.Stdout) //인코드 선언
// 	if err := encoder.Encode(p); err != nil { //인코드 실행
// 		panic(err)
// 	}
	
// 	fmt.Println(p)
// 	//json 문자열로 반환
// 	jdata, _ := json.Marshal(p)
// 	fmt.Println(string(jdata))
// }



func main() {
	p := Person{Name: "codz", Age: 23, Pnum: "01098765432"}
	//파일 오픈 - read/write 권한 부여 옵션
	file, _ := os.OpenFile("./person.json", os.O_CREATE|os.O_WRONLY, 0644)
	
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.Encode(p) //파일 출력


	file, _ := os.OpenFile("./person.json", os.O_RDONLY, 0644)
	decoder := json.NewDecoder(file) //decoder 선언

	// 스트림에서 읽은 데이터를 자료구조로 변환합니다.
	var p2 Person
	if err := decoder.Decode(&p2); err != nil && err != io.EOF {
	}
		fmt.Println(p2)
	
}