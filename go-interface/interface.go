// package main

// import "fmt"

// // Stringer라는 interface를 만들어 사용해보자

// type Stringer interface {
// 	String() string

// }

// type Student struct {
// 	Name string
// 	Age int
// }

// // Student의 String() 메서드 정의
// func(s Student) String() string {
// 	// 문자열을 만들어서 리턴
// 	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해.", s.Age, s.Name)
// }

// func main()  {
// 	student := Student{"철수", 12} // Student 타입 객체 생성
// 	var stringer Stringer // Stringer타입의 stringer 생성

// 	// student타입은 String()메서드를 포함하고 있기 때문에 stringer 값으로 student를 대입할 수 있다.
// 	stringer = student // stringer에 student 대입

// 	// fmt.Println(student.String())
// 	// return 된 문자열을 출력
// 	fmt.Printf("%s\n", stringer.String())
// }