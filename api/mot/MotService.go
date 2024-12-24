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

	// Rechercher les mots
	cursor, err := collection.Find(ctx, bson.D{{"first_letter", firstLetter}})
	if err != nil {
		return nil, err
	}
	// Récupérer les mots, s'il y a une erreur, la renvoyer
	if err = cursor.All(ctx, &mots); err != nil {
		return nil, err
	}

	return mots, nil
}


func GetMotFirstLetter(firstLetter string) (*Mot, error) {
	if utf8.RuneCountInString(firstLetter) != 1 {
		return nil, errors.New(InvalidFirstLetter)
	}

	collection := db.GetCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	matchStage := bson.D{{"$match", bson.D{{"first_letter", firstLetter}}}}
	sampleStage := bson.D{{"$sample", bson.D{{"size", 1}}}}

	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{matchStage, sampleStage})
	if err != nil {
		return nil, err
	}

	// Récupérer les mots
	var mots []Mot
	if err = cursor.All(ctx, &mots); err != nil {
		return nil, err
	}

	// Vérifier si aucun mot n'a été trouvé
	if len(mots) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return &mots[0], nil
}

func GetMotLength(length int) (*Mot, error) {
	collection := db.GetCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	matchStage := bson.D{{"$match", bson.D{{"length", length}}}}
	sampleStage := bson.D{{"$sample", bson.D{{"size", 1}}}}

	// Rechercher un mot de la longueur spécifiée
	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{matchStage, sampleStage})
	if err != nil {
		return nil, err
	}
	// Récupérer le mot
	var mots []Mot
	if err = cursor.All(context.TODO(), &mots); err != nil {
		return nil, err
	}

	// Si aucun mot n'a été trouvé, renvoyer une erreur
	if len(mots) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return &mots[0], nil
}