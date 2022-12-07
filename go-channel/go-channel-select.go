// package main

import (
	"fmt"
	"sync"
	"time"
)


func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for { // for문을 사용해서 반복하기 때문에 quit이 도착할 때까지 종료되지 않는다.
		select {
		case n := <-ch:
			fmt.Printf("Square: %d\n", n * n)
			time.Sleep(time.Second)
		case <- quit: // quit채널에 true라는 값이 전달되면 아래 종료 절차를 밟는다.
			wg.Done()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go square(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	// for문이 완료된 이후, 종료 시그널을 보낸다.
	quit <- true
	wg.Wait()
}