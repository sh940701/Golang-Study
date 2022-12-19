package main

import (
	"github.com/gin-gonic/gin"
	"lecture/go-swagger/controller"
	"lecture/go-swagger/router"
	"lecture/go-swagger/model"
	"net/http"
	"time"
)

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

func main() {
	md, _ := model.NewModel()
	ctl := controller.NewCTL(md)
	router, _ := router.NewRouter(ctl)

	mapi := http.Server {
		Addr : ":8080",
		Handler : router.Routing(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	mapi.ListenAndServe()
}