// package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	// 잔고가 0원 아래로 내려가면 panic 호출
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}

	// 아래 단계는 2 step으로 나뉜다. 값을 읽어오고, 여기에 1000원을 추가하는 것
	// 이 때 두 routine이 account.Balance == 0 으로 읽어온 후 둘 다 1000을 추가하면
	// 두 번 1000을 추가했음에도 불구하고 결과값은 그대로 1000이 된다.
	// 결과적으로 하위에서 두 번 -1000을 진행하게 되면 panic이 발생하는 것이다.
	account.Balance += 1000 // 잔고에 1000원 추가
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func main() {
	var wg sync.WaitGroup

	// 잔고가 0원인 계정 생성
	account := &Account{0}
	wg.Add(50)
	for i := 0; i < 50; i++ {
		go func ()  {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

/*
	panic: Balance should not be negative value: -1000

	goroutine 19 [running]:
	main.DepositAndWithdraw(0xc00001c0c8?)
			/Users/sunghyun/go/src/go-routine/go-routine-warning.go:16 +0xa8
	main.main.func1()
			/Users/sunghyun/go/src/go-routine/go-routine-warning.go:33 +0x25
	created by main.main
			/Users/sunghyun/go/src/go-routine/go-routine-warning.go:31 +0x53
	exit status 2
*/