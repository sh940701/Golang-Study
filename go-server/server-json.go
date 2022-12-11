package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// get 요청을 받으면 학생 데이터를 JSON 데이터로 반환하는 서버를 만들어보자.
type Student struct {
	Name string
	Age int
	Score int
}

// 아하 mux도 Handler중 한 종류구나
func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/student", StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	student := Student{"sh", 27, 9999}
	data, _ := json.Marshal(student) // student 객체를 JSON포맷으로 변환한다.
	// data는 []byte 타입 데이터가 된다.

	// response의 header에 데이터가 json타입임을 명시해준다.
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK) // 이건 꼭 써줘야하나? ㅇㅅㅇ .. ... . .....
	// fmt.Fprint(w, data) // 결과를 전송해준다. -> 이대로 보내면 각 byte가
	// Decimal 아스키코드에 해당하는 숫자로 전송이 된다.
	fmt.Fprint(w, string(data)) // string타입으로 변환하여 데이터 전송
}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}