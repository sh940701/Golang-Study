// package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body string
	Tire string
	Color string
	id int
}

// 전역적으로 사용할 WaitGroup을 만들어준다.
var wg sync.WaitGroup
var startTime = time.Now()

// 차체를 만드는 함수
func MakeBody (tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(time.Second * 10)
	id := 0
	for {
		select {
		case <- tick: // tick 채널의 데이터가 올 때마다 생산 시작
			car := new(Car)
			car.id = id
			id++
			car.Body = "Sports car"
			fmt.Printf("id %d making Body\n", car.id)
			// tire 채널로 car 객체를 보내준다.
			tireCh <- car

		// 종료시 tire채널을 닫고 고루틴을 종료한다.
		case <- after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

// 타이어를 설치하는 함수
func InstallTire (tireCh, paintCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		fmt.Printf("id %d installing tire\n", car.id)
		// 타이어 설치 작업 후, paint채널로 car 객체를 보내준다.
		car.Tire = "Winter tire"
		paintCh <- car
	}
	// tire채널이 close되면 아래로 내려온다.
	wg.Done() // 고루틴 종료
	close(paintCh) // 페인트 채널 닫음
}

// 페인트를 칠하는 함수
func PaintCar(paintCh chan *Car) {
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "Red"
		duration := time.Now().Sub(startTime) // 경과시간
		fmt.Printf("%d번 차량 생산 완료. 경과시간: %f초\n", car.id, duration.Seconds())
	}
	wg.Done()
}

func main() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("Start Factory\n")

	wg.Add(3)
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("Close the factory")
}
/*
	Start Factory
	id 0 making Body

	id 1 making Body
	id 0 installing tire

	id 2 making Body
	id 1 installing tire
	0번 차량 생산 완료. 경과시간: 3.007603초

	id 3 making Body
	id 2 installing tire
	1번 차량 생산 완료. 경과시간: 4.008226초

	id 4 making Body
	id 3 installing tire
	2번 차량 생산 완료. 경과시간: 5.008610초

	id 5 making Body
	id 4 installing tire
	3번 차량 생산 완료. 경과시간: 6.009700초

	id 6 making Body
	id 5 installing tire
	4번 차량 생산 완료. 경과시간: 7.009819초

	id 7 making Body
	id 6 installing tire
	5번 차량 생산 완료. 경과시간: 8.010629초

	id 8 making Body
	id 7 installing tire
	6번 차량 생산 완료. 경과시간: 9.011316초

	id 9 making Body
	id 8 installing tire
	7번 차량 생산 완료. 경과시간: 10.013791초

	id 9 installing tire
	8번 차량 생산 완료. 경과시간: 11.014890초

	9번 차량 생산 완료. 경과시간: 12.016290초
	
	Close the factory
*/