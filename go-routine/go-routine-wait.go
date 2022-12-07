// package main

import (
	"fmt"
	"sync"
)

// sync 패키지의 WaitGroup 객체를 사용한다.
var wg sync.WaitGroup

func SumAtoB(a, b int)  {
	sum :=0
	for i := a; i < b; i++ {
		sum += i
	}
	fmt.Printf("%d부터 %d까지 합계는 %d입니다.", a, b, sum)
	fmt.Println("")
	// 작업이 끝나면, 끝났다는 이벤트를 넘긴다.
	wg.Done()
}

func main() {
	// 몇개의 고루틴을 사용할지 선언해준다. 루틴 내에서 wg.Done이 호출되면
	// Add()로 선언되었던 개수가 1 감소한다.
	wg.Add(3)
	for i := 0; i < 100; i++ {
		go SumAtoB(1, 100000000000)
	}

	// 모든 고루틴의 작업이 종료될 때까지 기다린다.
	wg.Wait()

	fmt.Println("모든 작업이 완료되었습니다.")
}