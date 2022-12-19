package controller

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
	"io/ioutil"
	"lecture/go-swagger/model"
	"fmt"
)

type Controller struct {
	md *model.Model
}

func NewCTL(md *model.Model) (*Controller) {
	r := &Controller{md : md}
	return r
}

// GetOK godoc
// @Summary call GetOK, return ok by json.
// @Description api test를 위한 기능.
// @name GetOK
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /acc/v01/ok [get]
// @Success 200 {object} Controller
func (p *Controller) GetOK(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "ok"})
}


// CtlGetByName godoc
// @Summary call CtlGetByName with path parameter "name", return object by json.
// @Description 이름으로 원하는 person의 데이터를 찾는 기능
// @name CtlGetByName
// @Accept  json
// @Produce  json
// @Param name path string true "name"
// @Router /acc/v01/person/{name} [get]
// @Success 200 {object} model.Person
func (p *Controller) CtlGetByName(c *gin.Context) {
	name := c.Param("name")
	fmt.Println(name)
	result := p.md.GetInfoByNameFromDB(name)

	c.JSON(200, gin.H{"msg" : "ok", "result" : result})
}


// CtlAddPerson godoc
// @Summary call CtlAddPerson body "name", "age", "pnum", return object by json.
// @Description Person 데이터를 추가하는 기능
// @name CtlAddPerson
// @Accept  json
// @Produce  json
// @Param person body model.Person true "person"
// @Router /acc/v01/person/add [post]
// @Success 200 {object} model.Person
func (p *Controller) CtlAddPerson(c *gin.Context) {
	result := p.md.AddPersonFromDB(c)
	if result == nil {
		c.JSON(200, gin.H{"result" : "Fail"})
	} else {
		c.JSON(200, gin.H{"result" : "ok", "ObjectId" : result})
	}
}


// CtlUpdatePerson godoc
// @Summary call CtlUpdatePerson body "name", "age" return objectId
// @Description Person 데이터를 업데이트하는 기능
// @name CtlUpdatePerson
// @Accept  json
// @Produce  json
// @Param name body model.Person true "data"
// @Router /acc/v01/person/update [put]
// @Success 200 {object} string
func (p *Controller) CtlUpdatePerson(c *gin.Context) {
	body := c.Request.Body
	byteData, err := ioutil.ReadAll(body)

	if err != nil {
		panic(err)
	}

	var stringData map[string]interface{}
	json.Unmarshal(byteData, &stringData)

	name := stringData["name"]
	age := stringData["age"]

	result := p.md.UpdatePersonFromDB(name.(string), age.(float64))

	c.JSON(200, gin.H{"result" : "ok", "Before" : result[0], "After" : result[1]})
}