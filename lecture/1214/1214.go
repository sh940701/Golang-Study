package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func ping(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "world")
// }

// func main() {
// 	http.HandleFunc("/hello", ping)

// 	http.ListenAndServe("0.0.0.0:8080", nil)
// }

// import (
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()
// 	r.GET("/recv", func (c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"result" : "ok",
// 		})
// 	})
// 	r.Run()
// }

// import (
// 	"fmt"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// // 아하 uri binding을 할 때는 객체의 멤버변수가 대문자로 시작해야 하는구나
// // uri 바인딩을 위한 객체
// type Person struct {
// 	Name string `uri:"name" binding:"required"`
// 	Age int `uri:"age" binding:"required"`
// }

// // body 바인딩을 위한 객체
// type Student struct {
// 	Name string `json:"name" xml:"name" form:"name" binding:"required"`
// 	Grade int `json:"grade" xml:"grade" form:"grade" binding:"required"`
// }

// func get (c *gin.Context) {
// 	person := &Person{}
// 	if err := c.ShouldBindUri(&person); err != nil {
// 		c.JSON(400, gin.H{"msg" : err})
// 		return
// 	}
// 	fmt.Println(person)
// 	c.JSON(200, gin.H{"info" : "get " + person.Name + strconv.Itoa(person.Age), "msg" : "uri binding"})
// }

// func post (c *gin.Context) {
// 	student := &Student{}
// 	if err := c.ShouldBind(&student); err != nil {
// 		c.JSON(400, gin.H{"msg" : err})
// 		return
// 	}
// 	fmt.Println(student)
// 	c.JSON(200, gin.H{"info" : "post " + student.Name + strconv.Itoa(student.Grade), "msg" : "body binding"})
// }

// func main() {
// 	r := gin.Default()
// 	r.GET("/:name/:age", get)
// 	r.POST("/", post)
// 	r.Run()
// }

// func main() {
// 	router := gin.Default()

// 	v1 := router.Group("/v1") {
// 		v1.POST("/login", ex1)
// 		v1.PUT("/submit", ex2)
// 		v1.GET("/read", ex3)
// 	}

// 	v2 := router.Group("/v2") {
// 		v1.POST("/login", ex1)
// 		v1.PUT("/submit", ex2)
// 		v1.GET("/read", ex3)
// 	}

// 	router.Run(":8080")
// }

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

// user 정의 미들웨어

// 크로스 도메인 관용문구
//cross domain을 위해 사용
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//허용할 header 타입에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//허용할 method에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// 인증정보를 받아 처리해주는 미들웨어
func liteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c == nil {
			c.Abort()
		}
		auth := c.GetHeader("Authorization")
		fmt.Println("Authrization-word", auth)

		c.Next()
	}
}

// 라우팅
func index() *gin.Engine {
	// gin.Default()에는 기본적으로 Logger등의 기능이 있다.
	// gin.New()로 라우터를 생성해주면 미들웨어를 선택적으로 사용할 수 있다.
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORS())

	// 해당 그룹 라우터는 liteAuth() 미들웨어 사용
	route := r.Group("/route/v01", liteAuth())
	{
		route.POST("/post", post)
		route.PUT("/put", put)
		route.GET("/get", get)
		// group내의 group
		auth := route.Group("/auth")
		auth.GET("/nest", nest)
	}
	return r
}

func get (c *gin.Context) {
	c.JSON(200, gin.H{"msg" : "/route/v01/get"})
}

func post (c *gin.Context) {
	c.JSON(200, gin.H{"msg" : "/route/v01/post"})
}

func put (c *gin.Context) {
	c.JSON(200, gin.H{"msg" : "/route/v01/put"})
}

func nest (c *gin.Context) {
	c.JSON(200, gin.H{"msg" : "/route/v01/auth/nest"})
}

func main() {
	mapi := &http.Server{
		Addr: ":8080",
		Handler: index(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// 5초 지연 후 system call을 5초 지연해서 그 시간 이후에 꺼지게 하는건 너무 신기하다.
	// 1. 일단 stopSig라는 채널을 만든다.(그 안에 뭐가 들어간건진 아직 모르겠다..)
	// 2. 일단 control + c를 누르면 stopSig에 메시지가 전달되는 것 같다.
	// 3. 그럼 그 메시지가 메인함수에 들어오고, g.Engine객체의 Shutdown을 실행하게 되는데
	// 그 때 실행(종료하는 과정)을 context에게 맡긴다. context는 5초동안 지연후 처리하게 되고
	// 이로 인해 5초 뒤에 프로그램이 종료되는 것이다.(?) 이상 내 뇌피셜

	stopSig := make(chan os.Signal) //chan 선언
	// 해당 chan 핸들링 선언, SIGINT, SIGTERM에 대한 메세지 notify
	signal.Notify(stopSig, syscall.SIGINT, syscall.SIGTERM) 
	<-stopSig //메세지 등록

	// 해당 context 타임아웃 설정, 5초후 server stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := mapi.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
		case <-ctx.Done():
			fmt.Println("timeout 5 seconds.")
	}
	fmt.Println("Server stop")

	mapi.ListenAndServe()
}