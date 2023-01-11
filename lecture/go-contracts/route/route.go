package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ctl "lecture/go-contracts/controller"
)

type Router struct {
	ct *ctl.Controller
}

func NewRouter(ctl *ctl.Controller) (*Router, error) {
	r := &Router{ct : ctl}

	return r, nil
}

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

func liteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
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
	r := gin.New()
	
	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	r.Use(CORS())
	// root에서 /person 경로를 group화 해줌
	pRoute := r.Group("/person", liteAuth())
	{
		pRoute.GET("/symbol", p.ct.GetSymbol)
		pRoute.GET("/amount", p.ct.GetAmount)
		pRoute.POST("sendcoin", p.ct.SendCoin)
		pRoute.POST("/sendcoinbypk", p.ct.SendCoinByPk)
		pRoute.POST("/sendtoken", p.ct.SendToken)
		pRoute.POST("/sendtokenbypk", p.ct.SendTokenByPk)
	}
	return r;
}