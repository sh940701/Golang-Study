package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 핸들러 등록
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Print함수는 표준 출력 스트림으로 출력이 고정되는 반면
		// Fprint는 지정해준 출력 스트림에 값을 쓰는 함수이다.
		// http.ResponseWriter 타입에 값을 쓰면 HTTP 응답으로 전송된다.
		fmt.Fprint(w, "Hello, World")
	})

	// ListenAndServe함수에는 첫 번째 인자로 포트 번호, 두 번째 인자로 핸들러 인스턴스를
	// 넣어주는데, nil을 넣어주면 DefaultServeMux를 사용한다.
	// DefaultServeMux는 http.HandleFunc()함수를 호출해 등록된 핸들러들을 사용한다.
	http.ListenAndServe(":3000", nil)
}