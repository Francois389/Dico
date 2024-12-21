package mot

import (
	"Dico/db"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
	"unicode/utf8"
)

func GetMotsFirstLetter(firstLetter string) ([]Mot, error) {
	if utf8.RuneCountInString(firstLetter) != 1 {
		return nil, errors.New(InvalidFirstLetter)
	}

	var mots []Mot
	collection := db.GetCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, _ := collection.Find(ctx, bson.D{{"first_letter", firstLetter}})
	_ = cursor.All(context.TODO(), &mots)

	return mots, nil
}


func GetMotFirstLetter(firstLetter string) (Mot, error) {
	if utf8.RuneCountInString(firstLetter) != 1 {
		return NewMot("nil"), errors.New(InvalidFirstLetter)
	}

	var mots []Mot
	collection := db.GetCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	matchStage := bson.D{{"$match", bson.D{{"first_letter", firstLetter}}}}
	sampleStage := bson.D{{"$sample", bson.D{{"size", 1}}}}
	cursor, _ := collection.Aggregate(ctx, mongo.Pipeline{matchStage, sampleStage})
	_ = cursor.All(context.TODO(), &mots)

	var err error
	var resultat Mot

	if len(mots) == 0 {
		err = mongo.ErrNoDocuments
	} else {
		resultat = mots[0]
	}

	return resultat, err
}

func GetMotLength(length int) (Mot, error) {
	var mots []Mot
	collection := db.GetCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	matchStage := bson.D{{"$match", bson.D{{"length", length}}}}
	sampleStage := bson.D{{"$sample", bson.D{{"size", 1}}}}
	cursor, _ := collection.Aggregate(ctx, mongo.Pipeline{matchStage, sampleStage})
	_ = cursor.All(context.TODO(), &mots)

	var err error
	var resultat Mot

	if len(mots) == 0 {
		err = mongo.ErrNoDocuments
	} else {
		resultat = mots[0]
	}

	return resultat, err
}