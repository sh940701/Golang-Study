// package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// query 표현 방식
// http:://localhost:3000"?key=value&key=value"

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query() // 쿼리 인수가 담김
	name := values.Get("name") // 쿼리 인수 중 특정 키값이 있는지 확인
	if name == "" {
		name = "World"
	}

	id, _ := strconv.Atoi(values.Get("id")) // id 값을 가져와서 int 타입으로 변환
	// 이 때 id값이 없으면(string이 ""이면) 반환값은 숫자 0이다.
	fmt.Fprintf(w, "Hello %s! id:%d", name, id)
}

func main() {
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":3000", nil)
}