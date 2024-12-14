package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"sync"
	"time"
)

type Singleton struct {
	// Add fields here
	collection *mongo.Collection
	client     *mongo.Client
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

func getClient(ctx context.Context, url string) (*mongo.Client, error) {
	// Configurer un timeout de connexion
	clientOptions := options.Client().ApplyURI(url).SetConnectTimeout(10 * time.Second)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, fmt.Errorf("erreur de connexion à la base de données : %v", err)
	}

	// Vérifier que la connexion est établie
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("impossible de ping la base de données : %v", err)
	}

	fmt.Println("Connecté à la base de données")
	return client, nil
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

func Init(url, databaseName, collectionName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	instance := GetInstance()

	client, err := getClient(ctx, url)
	if err != nil {
		return fmt.Errorf("erreur lors de l'initialisation du client : %v", err)
	}

	instance.client = client
	database := getDatabase(client, databaseName)

	instance.SetCollection(getCollection(database, collectionName))
	return nil
}

func Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if instance == nil || instance.client == nil {
		return nil
	}

	err := instance.client.Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("erreur lors de la déconnexion : %v", err)
	}
	return nil
}
