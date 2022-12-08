package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	// context에 값을 추가한다. 여기서 key는 "number", value는 9이다.
	ctx := context.WithValue(context.Background(), "number", 9)
	go square(ctx)

	wg.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Printf("Square: %d\n", n * n) 
	}

	wg.Done()
}

/*
	취소도 되고 값도 설정하는 컨텍스트는 어떻게 만들까?
	-> 컨텍스트 객체를 만들어서 원하는 대로 wrapping해주면 된다.
	
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "number", 9)
	ctx = context.WithValue(ctx, "keyword", "Lilly")
*/