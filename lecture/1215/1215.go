package main

import (
	// "encoding/json"
	"fmt"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func main() {
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Mongo_URL))
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// db -> collection 단계적 접근
// 	db := client.Database("go-ready") // database 접속
// 	col0 := db.Collection("tPerson") // collection 접속

// 	// collection으로 바로 접근
// 	col1 := client.Database("go-ready").Collection("tPerson")

// 	// 한 DB 내 여러 collection 접근방식 -> db는 위에 선언되어있음
// 	colPerson := db.Collection("tPerson")
// 	colStudent := db.Collection("tStudent")
// 	colWoman := db.Collection("tWoman")
// }

// func main() {
// 	// mongodb password가 설정되있는경우
// 	credential := options.Credential{
// 		Username: "codz",
// 		Password: "states",
// 	}
// 	connect := func(dataSource string) (*mongo.Client, error) {
// 		if client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dataSource).SetAuth(credential)); err != nil {
// 			return nil, err
// 		} else if err = client.Ping(context.Background(), nil); err != nil { // 커넥션 유지를 위한 Ping
// 			return nil, err
// 		} else {
// 			return client, nil
// 		}
// 	}

// 	client, err := connect(Mongo_URL)
// }

// C
// const Mongo_URL = "mongodb://127.0.0.1:27017"

// type Person struct {
// 	Name string `bson:name`
// 	Age int `bson:age`
// 	Pnum string `bson:pnum,omitempty`
// }

// func main() {
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Mongo_URL))
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	db := client.Database("go-ready")
// 	collection := db.Collection("tPerson")

// 	newPerson := Person{Name: "inTest4", Age: 45}

// 	result, err := collection.InsertOne(context.TODO(), newPerson)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("Document inserted with ID: %s\n", result)
// }



// import (
// 	"encoding/json"
// 	"fmt"
// 	"context"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// // R - findOne
// const Mongo_URL = "mongodb://127.0.0.1:27017"

// type Person struct {
// 	Name string `bson:name`
// 	Age int `bson:age`
// 	Pnum string `bson:pnum,omitempty`
// }

// func main() {
// 	// mongodb와 연결
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Mongo_URL))
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// db접근
// 	db := client.Database("go-ready")
// 	// collection 접근
// 	collection := db.Collection("tPerson")

// 	// 해당 이름에 대한 접근
// 	strName := "inTest3"
// 	// strName := "codz"
// 	// filter에 해당하는 bson 객체
// 	filter := bson.D{{"name", strName}}
// 	// return 값을 넣을 bson 객체
// 	var result bson.M
// 	err = collection.FindOne(context.TODO(), filter).Decode(&result)

// 	if err == mongo.ErrNoDocuments {
// 		fmt.Printf("No document was found with the name %s\n", strName)
// 		return
// 	} else if err != nil {
// 		panic(err)
// 	}

// 	// rowdata 출력
// 	fmt.Println(result)

// 	// json으로 변경
// 	if jsonData, err := json.MarshalIndent(result, "", "     "); err != nil {
// 		panic(err)
// 	} else {
// 		// json화 된 data 출력
// 		fmt.Printf("%s\n", jsonData)
// 	}
// }


// const Mongo_URL = "mongodb://127.0.0.1:27017"

// type Person struct {
// 	Name string `bson:name`
// 	Age int `bson:age`
// 	Pnum string `bson:pnum,omitempty`
// }

// R - find
// func main() {
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Mongo_URL))
// 	db := client.Database("go-ready")
// 	collection := db.Collection("tPerson")

// 	// 결과를 담을 슬라이스
// 	results := []Person{}
// 	filter := bson.D{}
	
// 	// cursor 생성
// 	cursor, err := collection.Find(context.TODO(), filter)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer cursor.Close(context.TODO())

// 	// cursor.All과 cursor.Next, Decode를 사용한 결과값이 같다.
// 	// if err = cursor.All(context.TODO(), &results); err != nil {
// 	// 	panic(err)
// 	// }

// 	for cursor.Next(context.TODO()) {
// 		var p Person
// 		if err := cursor.Decode(&p); err != nil {
// 			panic(err)
// 		}
// 		results = append(results, p)
// 	}

// 	fmt.Println(results)
// }




// const Mongo_URL = "mongodb://127.0.0.1:27017"

// // U - updateOne

// func main() {
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Mongo_URL))
// 	db := client.Database("go-ready")
// 	collection := db.Collection("tPerson")

// 	filter := bson.D{{"name", "inTest3"}}
// 	update := bson.D{{"$set", bson.D{{"age", 10}}}}
// 	result, err := collection.UpdateOne(context.TODO(), filter, update)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("Documents Updated: %v\n", result.UpsertedCount)
// }




// const Mongo_URL = "mongodb://127.0.0.1:27017"

// // D - deleteOne

// func main() {
// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(Mongo_URL))

// 	db := client.Database("go-ready")
// 	collection := db.Collection("tPerson")

// 	option := bson.D{{"name", "inTest3"}}
// 	result, err := collection.DeleteOne(context.TODO(), option)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(result.DeletedCount)
// }



const Mongo_URL = "mongodb://127.0.0.1:27017"

func main() {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(Mongo_URL))
	db := client.Database("go-ready")
	collection := db.Collection("tPerson")

	strName := "codz"
	filter := bson.D{{"name", strName}}

	// 전체 data count 조회
	estCount, estCountErr := collection.EstimatedDocumentCount(context.TODO())
	if estCountErr != nil {
		panic(estCountErr)
	}
	fmt.Println("Total Document Count", estCount)

	// filter 조건에 맞는 data count 조회
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println("Filter Document Count", count)
}