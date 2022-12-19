package model

import (
	"log"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gin-gonic/gin"
)

type Model struct {
	client *mongo.Client
	collection *mongo.Collection
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
	} else if err = r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database("go-ready")
		r.collection = db.Collection("tPerson")
	}
	return r, nil
}

func (p *Model) GetInfoByNameFromDB(name string) *Person {
	person := &Person{}
	filter := bson.D{{Key: "name", Value: name}}
	err := p.collection.FindOne(
		context.TODO(),
		filter,
	).Decode(person)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		log.Fatal(err)
	}
	return person
}

func (p *Model) AddPersonFromDB(c *gin.Context) interface{} {
	newPerson := &Person{}
	if err := c.ShouldBindJSON(newPerson); err != nil {
		panic(err)
	}

	result, err := p.collection.InsertOne(context.TODO(), newPerson)

	if err != nil {
		panic(err)
	}

	return result.InsertedID
}

func (p *Model) UpdatePersonFromDB(name string, age float64) []*Person {
	var people []*Person

	people = append(people, p.GetInfoByNameFromDB(name))

	filter := bson.D{{Key: "name", Value: name}}
	update := bson.D{{"$set", bson.D{{Key: "age", Value: age}}}}

	_, err := p.collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}

	people = append(people, p.GetInfoByNameFromDB(name))

	return people
}