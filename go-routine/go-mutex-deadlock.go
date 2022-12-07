// package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	// var fork sync.Mutex -> 이렇게도 선언할 수 있당
	fork := &sync.Mutex{}
	spoon := &sync.Mutex{} // fork, spoon mutex 객체 생성

	go diningProblem("A", fork, spoon, "포크", "수저") // A는 포크 먼저
	go diningProblem("B", spoon, fork, "수저", "포크") // B는 수저 먼저
	wg.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 합니다.\n", name)
		first.Lock() // 첫 번째 뮤텍스를 획득 시도
		fmt.Printf("%s %s 획득\n", name, firstName)
		second.Lock() // 두 번째 뮤텍스를 획득 시도
		fmt.Printf("%s %s 획득\n", name, secondName)

		fmt.Printf("%s 밥을 먹습니다.\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock() // 뮤텍스 반납
		first.Unlock()
	}
}

/*
	B 밥을 먹으려 합니다.
	A 밥을 먹으려 합니다.
	B 수저 획득
	B 포크 획득
	B 밥을 먹습니다.
	A 포크 획득
	A 수저 획득
	A 밥을 먹습니다.
	B 밥을 먹으려 합니다.
	A 밥을 먹으려 합니다.
	A 포크 획득
	B 수저 획득
	fatal error: all goroutines are asleep - deadlock!
	
	서로 다른 go-routine에서 상호 공유하는 mutex에 대한 접근을 원할 때
	다른 routine이 이미 점유하고 있는 상황이라면 deadlock이 발생할 수 있다.
*/