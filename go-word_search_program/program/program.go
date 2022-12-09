package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 이 녀석은 단어가 포함되어있는 각 라인의 정보를 담는다.
type LineInfo struct {
	lineNo int
	line string
}

// 이 녀석은 파일 이름과, 포함되어있는 각 라인 구조체로 이루어진 배열을 담는다.
type FindInfo struct {
	filename string
	lines []LineInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 인수가 필요합니다.")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]
	var findInfos []FindInfo
	// findInfos := []FindInfo{} -> 이렇게 만들어줘도 된다.
	for _, path := range files {
		fmt.Println(path)
		// findInfos = FindWordInAllFiles(word, path) -> 이렇게 하면 마지막 한 녀석 것만 출력된다.
		// append는 슬라이스에 원소를 추가할 때 뿐 아니라, 슬라이스를 이어붙일 때도 사용할 수 있다.
		// append(slice, slice)와 같이 사용할 수 있는 것이다.
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}
	// 아래의 findInfo 에는 최종적으로 각 파일별로 word가 포함된 line에 대한 정보를 가지고 있는
	// 객체가 담겨지게 된다.
	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("=============================")
		// 각 라인별로 라인넘버와 문장을 출력
		for _, lineInfo := range findInfo.lines {
			fmt.Println("\t", lineInfo.lineNo, "\t", lineInfo.line)
		}
		fmt.Println("=============================")
		fmt.Println()
	}
}

func GetFileList(path string) ([]string, error) {
	// 이 녀석은 경로를 받아 경로에 있는 파일 리스트를 리턴해준다.
	return filepath.Glob(path)
}

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	// filelist에는 경로에 포함되어 있는 파일의 목록이 담겨있다.
	// 근데 지금 여기선 파일 하나만 담겨있다 ㅎㅎ
	filelist, err := GetFileList(path)
	if err != nil {
		fmt.Println("파일 경로가 잘못되었습니다. err: ", err, "path: ", path)
		return findInfos
	}

	// 경로에 포함되어있는 파일들을 대상으로 작업을 수행한다.
	for _, filename := range filelist {
		// 이 녀석은 파일 이름과 찾고자 하는 단어를 받아서, 파일 내에서 단어가 포함되어있는
		// LineInfo를 멤버변수로 가지고 있는 FindInfo 객체를 반환한다.
		// 그리고 파일별로 생성된 FindInfo 객체를 findInfos 슬라이스에 넣어준다.
		findInfos = append(findInfos, FindWordInFile(word, filename))
	}
	// 그리고 FindInfo 객체들로 이루어져 있는 findInfos 슬라이스를 리턴해준다.
	// 이 슬라이스는 main 함수의 findInfos 슬라이스와 합쳐지게 된다.
	return findInfos
}

func FindWordInFile(word, filename string) FindInfo {
	// findInfo 객체 생성 -> 이 객체는 파일당 하나씩 있으며, 내부적으로 
	// LineInfo로 이루어진 슬라이스에 단어가 포함된 line들을 가지게 된다.
	findInfo := FindInfo{filename, []LineInfo{}}
	// 파일 이름을 받아서 파일을 연다.
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", filename)
		return findInfo
	}
	// 함수 종료시에는 파일을 닫아줘야지
	defer file.Close()

	lineNo := 1
	// 파일을 string으로 읽는 scanner 객체 생성
	scanner := bufio.NewScanner(file)
	// 파일을 한 줄씩 읽는다.
	for scanner.Scan() {
		// 한 줄씩 읽어온 데이터를 Text()메서드를 사용해 변수에 저장한다.
		line := scanner.Text()
		// 위 line 변수에 word가 포함되어 있다면, 그 정보를 LineInfo 객체에 담아 
		// findInfo 객체의 lines 슬라이스에 담아준다.
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	return findInfo
}