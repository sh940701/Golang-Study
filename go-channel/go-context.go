// package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)


var wg sync.WaitGroup

func main() {
	wg.Add(1)
	// ctx, cancel := context.WithCancel(context.Background()) // context 생성

	// WithTimeout함수를 사용하면 두번째 인자(시간) 뒤에 Done()을 반환하는 ctx객체를 준다.
	ctx, _ := context.WithTimeout(context.Background(), time.Second * 3)
	go PrintEverySecond(ctx)
	// time.Sleep(5 * time.Second)
	// cancel() // 취소

	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <- ctx.Done(): // 취소 확인
			wg.Done()
			return
		
		case <- tick:
			fmt.Println("Tick")
		}
	}
}