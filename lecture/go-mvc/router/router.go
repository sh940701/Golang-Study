package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ctl "lecture/go-mvc/controller"
)

// 라우터 객체 - controller로 이루어져 있음
type Router struct {
	ct *ctl.Controller
}

// router를 만들어주는 함수 - controller를 받아서 만들어줌
func NewRouter(ctl *ctl.Controller) (*Router, error) {
	r := &Router{ct : ctl}

	return r, nil
}

// cross domain을 위해 만들어주는 미들웨어
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

// 사용자 인증을 위한 미들웨어
func liteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 인증 기능 구현
		if c == nil {
			c.Abort()
		}
		auth := c.GetHeader("Authorization")
		fmt.Println("Authorization-word: ", auth)

		c.Next()
	}
}

// 실제 라우팅 기능 함수
func (p *Router) Idx() *gin.Engine {
	// 구현해야 할 부분
	// 실제 gin.Engine객체를 만들어줌
	e := gin.New()
	
	e.Use(gin.Logger())
	e.Use(gin.Recovery())
	e.Use(CORS())
	account := e.Group("/acc/v01", liteAuth())
	{
		fmt.Println(account)
		account.GET("/ok", p.ct.GetOK) // controller 패키지의 실제 처리 함수
	}
	return e;
}