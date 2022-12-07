// package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// 채널 생성
	ch := make(chan int)
	// var ch chan string = make(chan string) -> 이렇게도 생성 가능

	// sync.WaitGroup에 하나의 고루틴 생성
	wg.Add(1)
	// 고루틴에서 square함수 생성 -> 인자값으로 WaitGroup과 channel 넘겨줌
	go square(&wg, ch)
	// 채널에 9 라는 데이터를 넣음
	ch <- 9
	// 서브 고루틴 작업이 끝날때까지 대기
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	// channel에서 데이터를 빼내어 n에 담아줌
	// 이 때, 채널에 데이터가 들어가는 것보다 고루틴이 먼저 실행되었기 때문에
	// 데이터가 들어올 때까지 기다리게 된다.
	n := <-ch
	// 1초간 대기
	// time.Sleep(time.Second)
	// channel에서 받아온 data의 제곱을 print
	fmt.Printf("Square: %d\n", n * n)
	// 고루틴 작업 끝
	wg.Done()
}