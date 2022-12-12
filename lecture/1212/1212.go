package main

import (
	"fmt"
	"time"
)

// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	//두개의 고루틴과, 한개의 함수가 say()를 호출, 순차적인 say()함수 실행
// 	go say("world")
// 	go say("codz")
// 	say("hello")
// }

// func main() {
// 	say("world")
// 	say("hello")
// }

// func counter(s string, n int, wg *sync.WaitGroup) {
//     for i := 0; i < n; i++ {
//         time.Sleep(time.Second)
//         fmt.Println(s, " : ", i)
//     }
// 	wg.Done()
// }

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(3)
// 	go counter("a", 10, &wg)

// 	go func() {
// 		counter("b", 10, &wg)
// 	}()

// 	go func(wg *sync.WaitGroup) {
// 		for i := 0; i < 10; i++ {
// 			time.Sleep(time.Second)
// 			fmt.Println("c : ", i)
// 		}
// 		wg.Done()
// 	}(&wg)

// 	wg.Wait()
// }




type Vector []float64

func (v Vector) DoSome(i, n int, u Vector, c chan int) {
    for ; i < n; i++ {
		time.Sleep(time.Second)
        v[i] += u[i]
		fmt.Println(v[i])
    }
	c <- 1
}

const numCPU = 4

func (v Vector) DoAll(u Vector) {
    c := make(chan int, numCPU)
    for i := 0; i < numCPU; i++ {
        go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
    }
    for i := 0; i < numCPU; i++ {
		<-c    // 한 태스크가 끝날 때까지 대기
    }
	close(c)
    // 모두 완료
}

func main() {
	var vu Vector = Vector{2, 5, 7}
	vu.DoAll(vu)
}


// ### 1) 실제 CPU 코어 수를 확인하여 설정 후 연산 처리 예)

// ```go
// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	runtime.GOMAXPROCS(runtime.NumCPU()) // CPU 개수를 구한 뒤 사용할 최대 CPU 개수 설정

// 	fmt.Println(runtime.GOMAXPROCS(0)) // 설정 값 출력

// 	s := "Hello, world!"

// 	for i := 0; i < 100; i++ {
// 		go func(n int) { // 익명 함수를 고루틴으로 실행
// 			fmt.Println(s, n)
// 		}(i)
// 	}

// 	fmt.Scanln()
// }
// ```