package main

// toml 예시
/*
import (
	"fmt"
	toml "lecture/go-toml/config"
	"flag"
)

func main() {
	result := toml.GetConfig("/Users/sunghyun/go/src/lecture/go-toml/config/config.toml")

	// print Server -> Port
	fmt.Println(result.Server.Port) // :8080

	// print Account -> db -> pass
	fmt.Println(result.DB["account"]["pass"]) // admin!@

	// print for work -> desc
	for _, value := range result.Work {
		fmt.Println(value.Desc) // log, exam
	}
}
*/


// flag 예시
/*
import (
	"fmt"
	"flag"
)

func main() {
	var port string

	// default값 선언
	flag.StringVar(&port, "port", "7070", "port to listen on")

	var conf string
	flag.StringVar(&conf, "config", "./config/config.toml", "config file to use")

	pMod := flag.String("mode", "debug", "service mode")
	flag.Parse()
	fmt.Println(port)
	fmt.Println(conf)
	fmt.Println(*pMod)
}
*/



// logger 예시
import (
	conf "lecture/go-toml/config"
	"time"
	"fmt"
	"net/http"
	"os"
	"lecture/go-toml/logger"
	"os/signal"
	"syscall"
	"context"
	rt "lecture/go-toml/route"
)

// var (
// 	g errgroup.Group
// )

func main() {
	// var configFlag = flag.String("config", "./conf/config.toml", "toml file to use for configuration")
	// flag.Parse()
	// cf := conf.NewConfig(*configFlag)

	cf := conf.GetConfig("/Users/sunghyun/go/src/lecture/go-toml-flag-logger/config/log-config.toml")

	// 로그 초기화
	if err := logger.InitLogger(cf); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	logger.Debug("ready server....")

	rt := rt.NewRouter()

		mapi := &http.Server{
			// Addr:           cf.Server.Port,
			Addr:           ":8080",
			Handler:        rt.Idx(),
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		// g.Go(func() error {
		// 	return mapi.ListenAndServe()
		// })

		mapi.ListenAndServe()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		logger.Warn("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := mapi.Shutdown(ctx); err != nil {
			logger.Error("Server Shutdown:", err)
		}

		select {
		case <-ctx.Done():
			logger.Info("timeout of 5 seconds.")
		}

		logger.Info("Server exiting")
	}

	// if err := g.Wait(); err != nil {
	// 	logger.Error(err)
	// }

