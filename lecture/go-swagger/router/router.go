package router

import (
	ctl "lecture/go-swagger/controller"
	"lecture/go-swagger/docs" //swagger에 의해 자동 생성된 package
	ginSwg "github.com/swaggo/gin-swagger"
	swgFiles "github.com/swaggo/files"
	"github.com/gin-gonic/gin"
)

type Router struct {
	ctl *ctl.Controller
}

func NewRouter(ctl *ctl.Controller) (*Router, error) {
	r := &Router{ctl : ctl}

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

func (p *Router) Routing() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORS())

	r.GET("/health")
	r.GET("/swagger/:any", ginSwg.WrapHandler(swgFiles.Handler))

	docs.SwaggerInfo.Host = "localhost:8080"

	accGroup := r.Group("/acc/v01")
	{
		accGroup.GET("/ok", p.ctl.GetOK)
		accGroup.GET("/person/:name", p.ctl.CtlGetByName)
		accGroup.POST("/person/add", p.ctl.CtlAddPerson)
		accGroup.PUT("/person/update", p.ctl.CtlUpdatePerson)
	}

	return r
}