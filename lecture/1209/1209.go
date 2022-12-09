package main

import (
	"fmt"
	"sync"
	"time"
)

// func main(){
// 	go counter("a", 10)

// 	go func() {
// 		counter("b", 10)
// 	}()

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			time.Sleep(time.Second)
// 			fmt.Println("c : ", i)
// 		}
// 	}()

// 	counter("d", 10)
// }

// func main() {
// 	// var wg sync.WaitGroup

// 	// wg.Add(2)

// 	ch := make(chan int, 2)
// 	ch <- 1 //입력
// 	ch <- 2 //입력
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

// var wg sync.WaitGroup

// func main() {
// 	wg.Add(1)
// 	ch := make(chan int)

// 	go func(ch chan int) {
// 		defer wg.Done()
// 		ch <- 2
// 		close(ch)
// 	}(ch)

// 	result := <- ch
// 	fmt.Println(result)

// 	wg.Wait()
// }

// func main() {
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	ch <- 3
// 	fmt.Println(ch)
// 	fmt.Println(ch)
// 	fmt.Println(ch)
// }

// func main() {
// 	ping := time.Tick(100 * time.Millisecond)
// 	pong := time.After(500 * time.Millisecond)
// 	for {
// 		select {
// 		case <-ping:
// 			fmt.Println("ping")
// 		case <-pong:
// 			fmt.Println("pong")
// 			return      //주석처리후 테스트
// 		// default:
// 		// 	fmt.Println("-")
// 			// time.Sleep(50 * time.Millisecond)
// 		}
// 	}
// }

// func counter(s string, n int) {
//     for i := 0; i < n; i++ {
//         time.Sleep(time.Second)
//         fmt.Println(s, " : ", i)
//     }
// }

// func main() {
// 	go counter("a", 10)

// 	go func() {
// 		counter("b", 10)
// 	}()

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			time.Sleep(time.Second)
// 			fmt.Println("c : ", i)
// 		}
// 	}()

// 	counter("d", 10)
// }

// func counter(i int) {
// 	fmt.Println(i)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 1; i <= 5; i++ {
// 		wg.Add(1)
// 		time.Sleep(time.Second)

// 		i := i
// 		go func() {
// 			defer wg.Done()
// 			counter(i)
// 		}()
// 	}

// 	wg.Wait()
// }

// func run() chan int {
// 	ch := make(chan int, 5)

// 	go func(ch chan int) {
// 		ch <- rand.Intn(100)
// 	}(ch)

// 	return ch
// }

// func main() {

// 	for i := 0; i <= 5; i++ {
// 		n := <- run()
// 		fmt.Println(n)
// 	}
// }

// var wg sync.WaitGroup

// func run(job string) {
// 	fmt.Println(job)
// 	time.Sleep(time.Second * 2)
// 	wg.Done()
// }

// func main() {
// 	var work = []string{"a", "b", "c", "d"}
// 	for i := 0; i <= 10; i++ {
// 		wg.Add(4)
// 		fmt.Println("work start", i)
// 		for _, job := range work {
// 			go run(job)
// 		}
// 		wg.Wait()
// 		fmt.Println("work end", i)
// 		fmt.Println()
// 	}
// }



func calc(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select { 
		case n := <- ch:
		fmt.Println(n + n)
		time.Sleep(time.Second * 2)
		case <- quit:
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
	go calc(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i
	}

	quit <- true

	wg.Wait()
}