package model

import (
	"context"
	// "encoding/json"
	// "fmt"
	// "net/http"

	// "github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (p *Model) GetPerson() []Person {
	// person 객체들을 가져오는 코드 구현
	var people []Person
	return people
}