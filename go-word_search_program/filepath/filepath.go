package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 실행 인수 읽고 파일 목록 가져오기
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다. ex) ex26.1 word filepath")
		return
	}
	// go run filepath.go word go* <- 이렇게 입력하면 word가 os.Args[1]에 들어간다.
	// files에는 path가 들어가야 한다는데 나는 path가 안들어가고 파일 목록이 들어간다;
	word := os.Args[1]
	files := os.Args[2:]
	fmt.Println("찾으려는 단어: ", word)
	PrintAllFiles(files)
}

func GetFileList(path string) ([]string, error) {
	// 파일 경로를 받아서, 해당하는 파일 목록을 가져오는 path/filepath 패키지의 함수
	return filepath.Glob(path)
}

func PrintAllFiles(files []string) {
	fmt.Println("찾으려는 파일 리스트")
	for _, path := range files {
		// 경로에 해당하는 filelist를 가져온다.
		filelist, err := GetFileList(path)
		if err != nil {
			fmt.Println("파일 경로가 잘못되었습니다. err: ", err, "path: ", path)
			return
		}
		// mac에서는 실행 인수로 파일 경로가 아닌 각 파일 리스트가 입력되었다. 
		// 그래서 여러번 출력이 된다. 킹바듬
		// fmt.Println("찾으려는 파일 리스트") -> 그래서 이걸 위로 올려줬다 ㅎㅎ
		for _, name := range filelist {
			fmt.Println(name)
		}
	}
}