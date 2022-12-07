// package main

import (
	"fmt"
	"sync"
	"time"
)


func square(wg *sync.WaitGroup, ch chan int) {
	// for range문을 사용하면 채널이 닫힐때까지 계속해서 데이터를 기다리게 된다.
	// 현재 코드에서는 채널을 닫지 않기 때문에 계속해서 기다리다가 Deadlock이 발생한다.
	for n := range ch {
		fmt.Printf("Square: %d\n", n * n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	// close 함수를 사용해서 채널을 닫아주었다. 이제 위에서 for range문이
	// 끝날 수 있고, wg.Done()이 호출되어 프로그램이 정상 종료 된다.
	close(ch)
	wg.Wait()
}