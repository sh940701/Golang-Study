// program1은 스레드 하나로 진행을 한다.
// program2에서는 go-routine과 channel을 이용해 분산처리를 해볼 것이다.

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type LineInfo struct {
	lineNo int
	line string
}

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

	for _, path := range files {
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}

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
	return filepath.Glob(path)
}

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}

	filelist, err := GetFileList(path)
	if err != nil {
		fmt.Println("파일 경로가 잘못되었습니다. err: ", err, "path: ", path)
		return findInfos
	}

	// 하나의 채널을 만들어준다. 이 채널은 FindInfo type 데이터를 다룬다.
	ch := make(chan FindInfo)
	cnt := len(filelist)
	recvCnt := 0

	// filelist에서 filename을 뽑아낸 후, 각 filename에 대해 goroutine으로
	// 내용에 단어가 포함되어 있는지 정보를 반환해주는 함수를 실행한다.
	// 이 때 채널도 인자로 함께 보낸다.
	for _, filename := range filelist {
		go FindWordInFile(word, filename, ch)
	}

	// 여기서 go-routine을 사용하는데 waitgroup을 통한 wait, done, add 등의 기능을
	// 사용하지 않아도 되는 이유는, 이미 아래 for문이 있기 때문이다.
	// wg.Wait()은, 서브 고루틴이 끝나기 전 메인 고루틴이 끝나지 않도록 해주는 것인데
	// 여기서는 이미 for문으로 서브 고루틴이 끝날때까지 기다리고 있기 때문에 해줄 필요가
	// 없는 것이다. -> 이상 뇌피셜

	// for문을 사용하여 채널에 데이터가 들어오길 기다린다.
	for findInfo := range ch {
		// 들어온 데이터를 findInfos라는 배열에 넣어준다.
		findInfos = append(findInfos, findInfo)
		recvCnt++

		// 만약, 파일의 수와 채널을 통해 받은 데이터의 수가 같으면
		// 파일 하나당 데이터 하나이기 때문에 for문을 종료해준다.
		if cnt == recvCnt {
			break
		}
	}

	return findInfos
}

func FindWordInFile(word, filename string, ch chan FindInfo) {
	findInfo := FindInfo{filename, []LineInfo{}}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", filename)
		// 직접 return해주는 것이 아니라 채널을 통해 데이터를 넘겨준다.
		ch <- findInfo
		return
	}

	defer file.Close()

	lineNo := 1

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	// 직접 return해주지 않고 채널을 통해 데이터를 넘겨준다.
	ch <- findInfo
}