package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 절대경로를 입력해줘야만 하다니...
	PrintFile("/Users/sunghyun/go/src/go-word_search_program/readline/hamlet.txt")
}

func PrintFile(filename string) {
	// os.Open을 이용해 파일을 연다.
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", filename)
		return
	}
	// 함수 종료 전 파일 닫기 필수!
	defer file.Close()

	// bufio.NewScanner()객체는 인자값으로부터 문자를 읽어오는 스캐너를 생성하여 반환한다.
	scanner := bufio.NewScanner(file)
	// Scan() 메서드는 "한 줄"을 읽어온다
	for scanner.Scan() {
		// 읽어온 한 줄의 데이터를 Text()메서드를 이용하여 변수에 담아준다.(여기선 바로 출력)
		fmt.Println(scanner.Text())
	}
}