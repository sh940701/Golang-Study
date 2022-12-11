// package main

import (
	"fmt"
	"net/http"
)

// ListenAndServe()함수의 두 번째 인수로 nil을 넣어서 DefaultServeMux를 사용하면
// http.HandleFunc()함수 같은 패키지 함수들을 이용해서 등록한 핸들러를 사용하기 때문에
// 다양한 기능을 추가하기 어려운 문제가 있다. 이를 ServeMux 인스턴스를 생성해서 해결할 수 있다.

func main() {
	mux := http.NewServeMux() // ServeMux 인스턴스 생성

	// mux객체에 귀속된 경로별 HandleFunc함수를 만들어줌
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar")
	})

	// ListenAndServe()함수의 두번째 인수로 ServeMux 인스턴스를 넣어서
	// DefaultServeMux가 아닌 새로 생성한 인스턴스를 사용하도록 한다.
	http.ListenAndServe(":3000", mux)
}