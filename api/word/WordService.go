package word

import (
	"Dico/db"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"time"
	"unicode/utf8"
)

func GetWordsFirstLetter(firstLetter string) ([]Word, error) {
	if utf8.RuneCountInString(firstLetter) != 1 {
		return nil, errors.New(InvalidFirstLetter)
	}

	var mots []Word
	collection := db.GetCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Find the words
	cursor, err := collection.Find(ctx, bson.D{{"first_letter", firstLetter}})
	if err != nil {
		return nil, err
	}
	// Get the words, if there is an error, return it
	if err = cursor.All(ctx, &mots); err != nil {
		return nil, err
	}

	return mots, nil
}

func GetWordFirstLetter(firstLetter string) (*Word, error) {
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

	// Retrieve the words
	var mots []Word
	if err = cursor.All(ctx, &mots); err != nil {
		return nil, err
	}

	// Check if no words were found
	if len(mots) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return &mots[0], nil
}

func GetWordLength(length int) (*Word, error) {
	collection := db.GetCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	matchStage := bson.D{{"$match", bson.D{{"length", length}}}}
	sampleStage := bson.D{{"$sample", bson.D{{"size", 1}}}}

	// Find a word of the specified length
	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{matchStage, sampleStage})
	if err != nil {
		return nil, err
	}
	// Retrieve the word
	var words []Word
	if err = cursor.All(context.TODO(), &words); err != nil {
		return nil, err
	}

	// If no words were found, return an error
	if len(words) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return &words[0], nil
}

/*
GetAnagrams retrieves a list of anagrams for the given word.
It returns a pointer to a slice of Word and an error if any occurs.
If no anagrams are found, it returns mongo.ErrNoDocuments.
*/
func GetAnagrams(mot string) ([]Word, error) {
	collection := db.GetCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Find the anagram list related to the given word (except the one given)
	cursor, err := collection.Find(ctx, bson.D{{"sorted_letter", sortLetter(mot)}, {"word", bson.D{{"$ne", mot}}}})
	if err != nil {
		return nil, err
	}

	// Retrieve the anagram list
	var mots []Word
	if err = cursor.All(context.TODO(), &mots); err != nil {
		return nil, err
	}

	// If no words were found, return an error
	if len(mots) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return mots, nil
}
