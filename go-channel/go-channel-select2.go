// package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	// time.Tick은 일정 시간을 주기로 신호를 보내주는 채널을 생성해 반환하는 함수
	// 일정 시간 간격으로 현재 시간을 나타내는 Time 객체를 반환한다.
	tick := time.Tick(time.Second) // 1초 간격 시그널

	// time.After는 현재 시간 이후로 일정 시간 경과 후에 신호를 보내는 채널을 생성해 반환하는 함수
	// 일정 시간 경과 후의 현재 시간을 나타내는 Time 객체를 반환한다.
	terminate := time.After(time.Second * 13)

	for {
		select {
		case <- tick: // 1초간격 action
			fmt.Println("Tick")

		case <- terminate: // 10초 뒤 action
			fmt.Println("Terminated!")
			wg.Done()
			return
		
		case n := <- ch:
			fmt.Printf("Square: %d\n", n * n)
			// 여기다 Sleep을 걸어놓으면, Tick이 1초마다 출력이 안되기도 한다.
			// 그 이유는 Sleep 함수가 아예 이 고루틴 자체를 재워버리기 때문이다.
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i ++ {
		ch <- i * 2
	}

	wg.Wait()
}