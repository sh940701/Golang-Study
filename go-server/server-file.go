package main

import "net/http"

// 루트경로에 있는 파일 읽어오기
// func main() {
// 	// 왜 그런지 내 컴퓨터에서는 기본 경로가 지금 폴더가 아니라 src 폴더로 되어있다.
// 	// 그렇기 때문에 경로를 이렇게 나열해줘야 한다.
// 	// 서버 시작 후 http://localhost:3000/golang.jpg <- 을 입력하면 사진이 출력된다.
// 	http.Handle("/", http.FileServer(http.Dir("go-server/static")))
// 	http.ListenAndServe(":3000", nil)
// }

// 다른 경로에 있는 파일 읽어오기
// 만약 http://localhost:3000/static/golang.jpg에 요청을 보내게 되면 어떻게 될까?
// 이 경우 static 경로가 추가되어서 웹서버에서 static/static/golang.jpg를 찾게 된다.
// http.Handle("/static/", http.FileServer(http.Dir("static")))
// 위와 같이 해줘도, 시스템 내부적으로 static/static/golang.jpg를 찾게된다.
// 이 때 경로를, 내가 원하는 경로로 핸들링 하는 방법을 알아보자
func main() {
	// http.SrtipPrefix()함수를 사용하여 URL에서 /static/을 제거해주면 된다.
	// 결국 http://localhost:3000/static/golang.jpg -> http://localhost:3000/golang.jpg
	// 위와 같이 변경된 접근으로 서버는 인식하는 것이다.
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("go-server/static"))))
	http.ListenAndServe(":3000", nil)

}