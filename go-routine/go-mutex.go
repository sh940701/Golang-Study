// package main

import (
	"fmt"
	"sync"
	"time"
)

// sync 패키지에서 mutex 객체를 사용
var mutex sync.Mutex

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	// mutex의 Lock을 걸어주어 현재 스코프를 임계구역으로 설정
	mutex.Lock()
	// 함수 종료시 Unlock해줄 것을 선언
	defer mutex.Unlock()
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func main() {
	// sync 패키지의 WaitGroup 객체를 사용
	var wg sync.WaitGroup
	account := &Account{0}

	// go-routine 10개를 사용할 것이라고 선언
	wg.Add(10)
	for i := 0; i < 10; i++ {
		// go-routine 실행 선언
		go func() {
			for {
				// DepositAndWithdraw() 함수 무한 반복 실행
				DepositAndWithdraw(account)
			}
			// 끝나면 Done() 함수를 호출하여 go-routine 회수
			// 근데 여기선 어차피 안 끝나는 함수라서 안해줘도 상관 X
			wg.Done()
		}()
	}

	// wg.Add(10)된 go-routine들이 끝날때까지 main thread가 대기
	wg.Wait()
}