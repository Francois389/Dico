package db

import (
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"sync"
)

type Singleton struct {
	// Add fields here
	collection *mongo.Collection
	client *mongo.Client
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func (s *Singleton) SetCollection(collection *mongo.Collection) {
	s.collection = collection
}

func (s *Singleton) getCollection() *mongo.Collection {
	return s.collection
}

func getClient(url string) *mongo.Client {
	client, err := mongo.Connect(options.Client().ApplyURI(url))
	if err != nil {
		fmt.Println("Error while connecting to the database")
		panic(err)
	}

	fmt.Println("Connected to the database")

	return client
}

func getDatabase(client *mongo.Client, databaseName string) *mongo.Database {
	return client.Database(databaseName)
}

func getCollection(database *mongo.Database, collectionName string) *mongo.Collection {
	return database.Collection(collectionName)
}

func GetCollection() *mongo.Collection {
	return GetInstance().getCollection()
}

func Init(url, databaseName, collectionName string) {
	instance := GetInstance()

	client := getClient(url)
	instance.client = client
	database := getDatabase(client, databaseName)

	instance.SetCollection(getCollection(database, collectionName))
}

func Close() {
	err := instance.client.Disconnect(nil)
	if err != nil {
		panic(err)
	}
}