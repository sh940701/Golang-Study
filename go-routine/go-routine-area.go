package main

import (
	"fmt"
	"sync"
	"time"
)

type Job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", j.index, j.index * j.index)
}

func main() {
	var jobList [10]Job
	// var jobList [10]SquareJob -> 이렇게 선언해도 되더라
	// jobList := make([]SquareJob, 10) -> 이렇게도 되더라

	// 10가지 작업 할당
	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
		// jobList[i] = SquareJob{i} -> 이렇게 선언해도 되더라
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobList[i]
		// job := &SquareJob{i} -> 이렇게 선언해도 되더라
		go func() { // 각 작업을 각각의 go-routine으로 실행
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}