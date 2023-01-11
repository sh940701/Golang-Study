package main

import (
	"fmt"
	"lecture/go-contracts/controller"
	"lecture/go-contracts/route"
	"net/http"
	"time"

	// "golang.org/x/sync/errgroup"
)

// var (
// 	g errgroup.Group
// )

func main() {
	// controller 객체 초기화
	if controller, err := controller.NewCTL(model); err != nil {
		fmt.Println(err)
		return
	// controller 객체를 포함하는 router 객체 초기화
	} else if router, err := router.NewRouter(controller); err != nil {
		fmt.Println(err)
		return
	} else {
		// 서버 오픈
		mapi := &http.Server{
			Addr: ":8080",
			Handler: router.Idx(),
			ReadTimeout: 5 * time.Second,
			WriteTimeout: 10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		mapi.ListenAndServe()
	}

}