
	// 왜 그런지 내 컴퓨터에서는 기본 경로가 지금 폴더가 아니라 src 폴더로 되어있다.
	// 그렇기 때문에 경로를 이렇게 나열해줘야 한다.
	// 서버 시작 후 http://localhost:3000/golang.jpg <- 을 입력하면 사진이 출력된다.
	http.Handle("/", http.FileServer(http.Dir("go-server/static")))
	http.ListenAndServe(":3000", nil)
}