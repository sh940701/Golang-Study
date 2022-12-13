package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 서버 그냥 시작해보기
// func setupRouter() *gin.Engine {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.String(200, "pong")
// 	})
// 	return r
// }

// func main() {
// 	r := setupRouter()
// 	r.Run(":8080")
// }

// 서버를 통해서 json파일 보내보기
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
			"message" : "pong",
		})
	})
	r.Run()
}