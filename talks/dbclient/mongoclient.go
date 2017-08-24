package dbclient

import (
	"fmt"
	"log"
	"time"

	"github.com/icrowley/fake"
	"github.com/mattstratton/blondie/talks/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// IMongoClient is an interface for Mongo
type IMongoClient interface {
	OpenMongoDB()
	QueryTalk(talkID string) (model.Talk, error)
	Seed()
}

// MongoClient defines a mongo session
type MongoClient struct {
	mongoDB *mgo.Session
}

// OpenMongoDB provides a connection to a mongoDB instance
func (mc *MongoClient) OpenMongoDB() {
	var err error
	mc.mongoDB, err = mgo.Dial("db") //TODO: Make this not hard-coded
	if err != nil {
		log.Fatal(err)
	}
}

// Seed creates a bunch of fake talks for fun
func (mc *MongoClient) Seed() {

	total := 100
	for i := 0; i < total; i++ {
		key := bson.NewObjectId()
		acc := model.Talk{
			ID:      key,
			Title:   fake.Sentence(),
			Summary: fake.Paragraphs(),
			When:    time.Now(),
		}

		// jsonBytes, _ := json.Marshal(acc)

		if err := mc.mongoDB.DB("blondie").C("talks").Insert(acc); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Seeded %v fake talks...\n", total)
}

func (mc *MongoClient) QueryTalk(talkID string) (model.Talk, error) {
	talk := model.Talk{}
	objectID := bson.ObjectIdHex(talkID)
	err := mc.mongoDB.DB("blondie").C("talks").Find(bson.M{"_id": objectID}).One(&talk)

	if err != nil {
		return model.Talk{}, err
	}

	return talk, nil

}
