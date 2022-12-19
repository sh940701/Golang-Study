package controller


import (
	"lecture/go-mvc/model"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"io/ioutil"
	"fmt"
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

// 전체 목록을 가져오는 controller 함수
func (p *Controller) CtlGetPeople(c *gin.Context) {
	result := p.md.GetPeopleFromDB()
	c.JSON(200, gin.H{"people" : result})
	// return
}

// 이름으로 data를 가져오는 controller 함수
func (p *Controller) CtlGetInfoByName(c *gin.Context) {
	param := c.Param("name")
	result := p.md.GetInfoByNameFromDB(param)
	// 리턴된 데이터가 nil이면 없다는 의미이므로 예외처리
	if result == nil {
		c.JSON(200, gin.H{"result" : "0 infos"})
		return
	}
	// 있다면 결과값 반환
	c.JSON(200, gin.H{"result" : result})
}

// pnum으로 data를 가져오는 controller 함수
func (p *Controller) CtlGetInfoByPnum(c *gin.Context) {
	param := c.Param("pnum")
	result := p.md.GetInfoByPnumFromDB(param)

	if result == nil {
		c.JSON(200, gin.H{"result" : "0 infos"})
		return
	}
	// 있다면 결과값 반환
	c.JSON(200, gin.H{"result" : result})
}

// person 데이터를 추가하는 controller함수
func (p *Controller) CtlAddPerson(c *gin.Context) {
	result := p.md.AddPersonFromDB(c)

	c.JSON(200, gin.H{"objectId": result})
}

// person 데이터를 삭제하는 controller 함수
func (p *Controller) CtlDeletePerson(c *gin.Context) {
	// body의 데이터를 가져오는 과정

	body := c.Request.Body // 여기서 body는 ReadCloser 객체이다.
	// stream 데이터 형식으로, 읽고 닫히는 것이다.(?ㅋㅋ 일단 여기까지)

	// Request.Body는 스트림 데이터로, 한 번 읽고나면 더이상 읽을 수 없다.
	// https://etloveguitar.tistory.com/99?category=902018

	// ioutil.ReadAll() 메서드로 stream데이터를 읽어오자
	data, _ := ioutil.ReadAll(body) // data의 타입은 []byte이다. -> 바이트배열로 출력됨
	fmt.Println("==========data0==========", data)
	fmt.Println("==========data0==========", data)

	var value map[string]interface{}
	// json.Unmarshal 메서드는 json 형식의 데이터를 []byte 타입으로 받아서
	// 두번째 인자의 지원하는 타입으로 변경해준다.
	json.Unmarshal(data, &value)

	// pnum값으로 삭제할 요소를 정할것이기 때문에 전달받은 body의 pnum값을 넘겨준다.
	pnum := value["pnum"]
	fmt.Println("==========value==========", value)
	result := p.md.DeletePersonFromDB(pnum.(string))

	var msg string
	if result == 0 {
		msg = "fail to delete"
	} else {
		msg = "success to delete"
	}

	c.JSON(200, gin.H{"msg": msg})
}

// person 데이터를 업데이트 하는 controller함수
func (p *Controller) CtlUpdatePerson(c *gin.Context) {
	// body는 ReadCloser 스트림 데이터
	body := c.Request.Body
	// 스트림 데이터를 읽어서 []byte 형식으로 반환
	data, _ := ioutil.ReadAll(body)
	// 결과를 담을 map 자료구조
	var value map[string]interface{}
	// json형식의 data([]byte)를 Unmarshal하며 map에 담아줌
	json.Unmarshal(data, &value)
	// map형식이 담긴 value의 "pnum", "age" value를 가져옴
	pnum := value["pnum"]
	age := value["changeAge"]

	fmt.Println("value======", value)

	result := p.md.UpdatePersonFromDB(pnum.(string), age.(float64))

	c.JSON(200, gin.H{"msg": "updateSucceed!", "before": result[0], "after" : result[1]})
}