package main

import (
	"Dico/db"
	"Dico/word"
	"bufio"
	"context"
	"flag"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
)

func main() {
	// Define a flag to clear existing data
	clearExisting := flag.Bool("clear", false, "Clear existing data before populating the database")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	db.Init("mongodb://localhost:27027", "dico-db", "mots")
	collection := db.GetCollection()
	defer db.Close()

	// Delete existing data if flag is set
	if *clearExisting {
		if _, err := collection.DeleteMany(ctx, bson.D{}); err != nil {
			log.Fatalf("Error when delete existing data: %v", err)
		}
		fmt.Println("Existing data has been deleted")
	}

	file, err := os.Open("mots.txt")
	if err != nil {
		log.Fatalf("Th file can't be opened: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	batchSize := 1000
	var wordsBatch []interface{}

	for scanner.Scan() {
		motTexte := cleanWord(scanner.Text())

		if isValidWord(motTexte) {
			wordsBatch = append(wordsBatch, word.NewWord(motTexte))

			if len(wordsBatch) >= batchSize {
				if !addWordsToCollection(collection, ctx, wordsBatch) {
					return
				}

				count += len(wordsBatch)
				fmt.Printf("Processed %d words\n", count)

				wordsBatch = nil
			}
		}
	}

	// Insert the remaining words
	if len(wordsBatch) > 0 {
		if !addWordsToCollection(collection, ctx, wordsBatch) {
			return
		}

		count += len(wordsBatch)
		fmt.Printf("Processed %d words\n", count)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d words have been added to the database.\n", count)
}

func addWordsToCollection(collection *mongo.Collection, ctx context.Context, wordsBatch []interface{}) bool {
	_, err := collection.InsertMany(ctx, wordsBatch)
	if err != nil {
		fmt.Printf("Error when adding words : %v", err)
		return false
	}
	return true
}

func cleanWord(word string) string {
	word = strings.ToLower(word)
	return strings.TrimSpace(word)
}

func isValidWord(word string) bool {
	for _, r := range word {
		if !unicode.IsLetter(r) {
			fmt.Printf("The word %s isn't valid\n", word)
			return false
		}
	}

	return true
}
