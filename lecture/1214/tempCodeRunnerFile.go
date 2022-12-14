import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 아하 uri binding을 할 때는 객체의 멤버변수가 대문자로 시작해야 하는구나
// uri 바인딩을 위한 객체
type Person struct {
	Name string `uri:"name" binding:"required"`
	Age int `uri:"age" binding:"required"`
}

// body 바인딩을 위한 객체
type Student struct {
	Name string `json:"name" xml:"name" form:"name" binding:"required"`
	Grade int `json:"grade" xml:"grade" form:"grade" binding:"required"`
}

func get (c *gin.Context) {
	person := &Person{}
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(400, gin.H{"msg" : err})
		return
	}
	fmt.Println(person)
	c.JSON(200, gin.H{"info" : "get " + person.Name + strconv.Itoa(person.Age), "msg" : "uri binding"})
}

func post (c *gin.Context) {
	student := &Student{}
	if err := c.ShouldBind(&student); err != nil {
		c.JSON(400, gin.H{"msg" : err})
		return
	}
	fmt.Println(student)
	c.JSON(200, gin.H{"info" : "post " + student.Name + strconv.Itoa(student.Grade), "msg" : "body binding"})
}

func main() {
	r := gin.Default()
	r.GET("/:name/:age", get)
	r.POST("/", post)
	r.Run()
}