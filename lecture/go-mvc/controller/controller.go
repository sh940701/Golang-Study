package controller


import (
	"lecture/go-mvc/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller 타입 선언
type Controller struct {
	// 내부적으로 model객체를 가지고있다.
	md *model.Model
}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{md : rep}
	return r, nil
}

func (p *Controller) RespOK(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, resp)
}

func (p *Controller) GetOK(c *gin.Context) {
	c.JSON(200, gin.H{"msg" : "ok"})
	return
}