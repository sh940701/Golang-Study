package model

import (
	"log"
	"context"
	// "encoding/json"
	// "fmt"
	// "net/http"
	// "reflect"
	// "github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gin-gonic/gin"
)

type Model struct {
	client *mongo.Client
	colPersons *mongo.Collection
}

type Person struct {
	Name string `bson:"name"`
	Age int `bson:"age"`
	Pnum string `bson:"pnum"`
}

func NewModel() (*Model, error) {
	r := &Model{}

	var err error
	mgUrl := "mongodb://127.0.0.1:27017"
	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl)); err != nil {
		return nil, err
	} else if err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database("go-ready")
		r.colPersons = db.Collection("tPerson")
	}
	return r, nil
}

// 새로운 Person 객체를 생성하는 함수
func (p *Model) NewPerson() *Person {
	var person Person
	return &person
}


// 전체 데이터를 가져오는 model 함수
func (p *Model) GetPeopleFromDB() []Person {
	// person 객체들을 가져오는 코드 구현
	var people []Person
	filter := bson.D{{}}
	
	cursor, err := p.colPersons.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	if err = cursor.All(context.TODO(), &people); err != nil {
		panic(err)
	}

	return people
}

// 이름을 받아 정보를 가져오는 model 함수
func (p *Model) GetInfoByNameFromDB(name string) *Person {
	person := &Person{}
	filter := bson.D{{Key: "name", Value: name}}
	// name을 가지고 findOne을 한다.
	err := p.colPersons.FindOne(
		context.TODO(),
		filter,
	).Decode(person)
	if err != nil {
		// name과 매칭되는 data가 없으면 nil을 리턴
		if err == mongo.ErrNoDocuments {
			return nil
		}
		log.Fatal(err)
	}
	// name과 매칭되는 data가 있으면 data를 리턴
	return person
}

// 번호를 받아 정보를 가져오는 model 함수
func (p *Model) GetInfoByPnumFromDB(pnum string) *Person {
	person := &Person{}
	filter := bson.D{{Key: "pnum", Value: pnum}}
	// pnum을 가지고 findOne을 한다.
	err := p.colPersons.FindOne(
		context.TODO(),
		filter,
	).Decode(&person)
	if err != nil {
		// name과 매칭되는 data가 없으면 nil을 리턴
		if err == mongo.ErrNoDocuments {
			return nil
		}
		log.Fatal(err)
	}
	// pnum과 매칭되는 data가 있으면 data를 리턴
	return person
}

// DB에 정보를 추가하는 model 함수
func (p *Model) AddPersonFromDB(c *gin.Context) interface{} {
	newPerson := Person{}
	// request의 payload에 있는 json 형식 데이터들을 Person객체 구조에 맞춰 binding해준다.
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		panic(err)
	}
	result, err := p.colPersons.InsertOne(context.TODO(), newPerson)
	if err != nil {
		panic(err)
	}
	return result.InsertedID
}

// DB에서 정보를 삭제하는 model 함수
func (p *Model) DeletePersonFromDB(pnum string) int64 {
	filter := bson.D{{Key: "pnum", Value: pnum}}
	result, err := p.colPersons.DeleteOne(context.TODO(), filter)
	if err != nil {
		return 0
	}
	return result.DeletedCount
}

func (p *Model) UpdatePersonFromDB(pnum string, age float64) []*Person {
	var people []*Person

	// 위에서 만든 model 함수 사용
	people = append(people, p.GetInfoByPnumFromDB(pnum))

	filter := bson.D{{Key: "pnum", Value: pnum}}
	update := bson.D{{"$set", bson.D{{Key: "age", Value: age}}}}

	// 업데이트 실행
	_, err := p.colPersons.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}

	// 업데이트 반영된 객체 가져오기
	people = append(people, p.GetInfoByPnumFromDB(pnum))

	return people
}