package mot

import (
	"API-Mot/db"
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
)

func GetMotsFirstLetter(firstLetter string) []Mot {
	var mots []Mot
	collection := db.GetCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, _ := collection.Find(ctx, bson.D{{"first_letter", firstLetter}})
	_ = cursor.All(context.TODO(), &mots)
	return mots
}


func GetMotFirstLetter(firstLetter string) Mot {
	var mots []Mot
	collection := db.GetCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	matchStage := bson.D{{"$match", bson.D{{"first_letter", firstLetter}}}}
	sampleStage := bson.D{{"$sample", bson.D{{"size", 1}}}}
	cursor, _ := collection.Aggregate(ctx, mongo.Pipeline{matchStage, sampleStage})
	_ = cursor.All(context.TODO(), &mots)
	return mots[0]
}