package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	Start = 1000
	Earn = 500
	Lose = 200
	WinGame = 5000
	LoseGame = 0
)

// bufio는 읽기 전용 객체를 만들어준다. 그 객체는 os.Stdin을 통해 받은 값으로 구성된다.
var stdin = bufio.NewReader(os.Stdin) 
// 입력값이 올바르지 않을 시 버퍼를 비워주는 에러 핸들링 함수
func InputIntValue() (int, error) {
	var n int
	// Scan으로 받은 값이 n의 타입과 다르면 err != nil 일 것이다.
	_, err := fmt.Scan(&n)
	if err != nil {
		// 입력값 에러 발생시 버퍼를 비워줌
		stdin.ReadString('\n')
	}
	return n, err
}

func main() {
	// 기본 값
	amount := Start

	for {
	fmt.Println("")
	// 랜덤 추출값
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(5) + 1
	
	fmt.Println("1 ~ 5의 숫자를 적어주세요")
	// 에러처리 함수를 실행한 후, 에러가 있을 시 for문 다시 시작
	input, err := InputIntValue()
	if err != nil || input < 1 || input > 5 {
		continue
	}
	
	if input == random {
		amount += Earn
		fmt.Println("맞추셨습니다!")
		fmt.Println("잔액: ", amount)
	} else {
		amount -= Lose
		fmt.Println("틀렸습니다!")
		fmt.Println("잔액: ", amount)
	}

	if amount >= WinGame {
		fmt.Println("축하합니다. 승리하셨습니다!")
		fmt.Println("잔액: ", amount)
		break
	} else if amount <= LoseGame {
		fmt.Println("패배하셨습니다.")
		fmt.Println("잔액: 0원")
		break
	}
}
}