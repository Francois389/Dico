package word

import (
	"api/db"
	"api/models"
	"context"
	"errors"
	"time"
	"unicode/utf8"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetWordsFirstLetter(firstLetter string) ([]models.Word, error) {
	if utf8.RuneCountInString(firstLetter) != 1 {
		return nil, errors.New(InvalidFirstLetter)
	}

	var mots []models.Word
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

func GetWordFirstLetter(firstLetter string) (*models.Word, error) {
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

	// Retrieve the Words
	var mots []models.Word
	if err = cursor.All(ctx, &mots); err != nil {
		return nil, err
	}

	// Check if no Words were found
	if len(mots) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return &mots[0], nil
}

func GetWordLength(length int) (*models.Word, error) {
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
	var words []models.Word
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
func GetAnagrams(mot string) ([]models.Word, error) {
	collection := db.GetCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Find the anagram list related to the given word (except the one given)
	cursor, err := collection.Find(ctx, bson.D{{"sorted_letter", models.SortLetter(mot)}, {"word", bson.D{{"$ne", mot}}}})
	if err != nil {
		return nil, err
	}

	// Retrieve the anagram list
	var mots []models.Word
	if err = cursor.All(context.TODO(), &mots); err != nil {
		return nil, err
	}

	// If no words were found, return an error
	if len(mots) == 0 {
		return nil, mongo.ErrNoDocuments
	}

	return mots, nil
}

// GetWordsBatch retrieves a batch of words for each letter in the given string.
func GetWordsBatch(letters string) []models.Word {
	// Find a random word starting with each letter in the given string
	var words []models.Word
	for _, letter := range letters {
		word, _ := GetWordFirstLetter(string(letter))
		if word != nil {
			words = append(words, *word)
		}
	}

	return words
}

func CheckWordExistence(word string) (bool, error) {
	collection := db.GetCollection()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, bson.D{{"word", word}})

	if err != nil {
		return false, err
	}
	return 0 < count, nil
}
