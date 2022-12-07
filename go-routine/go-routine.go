// package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c", v)
	}
	fmt.Println("")
}

func PrintNumbers() {
	for i:= 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d", i)
	}
	fmt.Println("")
}

func main() {
	// go 라는 명령어를 통해서, 서브 고루틴을 통해 연산을 시킬 수 있다.
	go PrintHangul()
	go PrintNumbers()

	// 서브 고루틴은, 메인 고루틴이 작동중일때만 정상동작한다.
	// 만약 메인 고루틴이 먼저 끝나버리면 서브 고루틴도 그대로 종료된다.
	time.Sleep(3 * time.Second)
}