package route

import (
	logger "lecture/go-toml/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Router struct {}

func NewRouter() *Router {
	r := &Router{}
	return r
}

func (p *Router) Idx() *gin.Engine {
	// 컨피그나 상황에 맞게 gin 모드 설정
	gin.SetMode(gin.ReleaseMode)
	// gin.SetMode(gin.DebugMode)

	e := gin.Default()
	// e.Use(gin.Logger())
	// e.Use(gin.Recovery())
  // 기존의 logger, recovery 대신 logger에서 선언한 미들웨어 사용
	e.Use(logger.GinLogger())
	e.Use(logger.GinRecovery(true))
	// e.Use(CORS())

	logger.Info("start server")
	e.GET("/health")

	// account := e.Group("acc/v01", liteAuth())
	account := e.Group("acc/v01")
	{
		fmt.Println(account)
	}

	return e
}