package main

import (
	"Dico/db"
	"Dico/mot"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
)

func main() {
	// Contexte avec timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()
	// Ouverture de la connexion à la base de données
	db.Init("mongodb://localhost:27027", "dico-db", "mots")
	collection := db.GetCollection()  // Adaptez selon votre méthode de connexion

	defer db.Close()

	// Ouvrir le fichier contenant les mots
	file, err := os.Open("mots.txt")
	if err != nil {
		log.Fatalf("Impossible d'ouvrir le fichier : %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	compteur := 0
	batchSize := 1000
	var wordsBatch []interface{}

	// Parcourir le fichier ligne par ligne
	for scanner.Scan() {
		motTexte := cleanWord(scanner.Text())

		if isValidWord(motTexte) {
			// Créer un nouveau mot
			nouveauMot := mot.NewMot(motTexte)  // Adaptez selon votre méthode de création

			// Ajouter le mot au lot
			wordsBatch = append(wordsBatch, nouveauMot)

			if len(wordsBatch) >= batchSize {
				_, err := collection.InsertMany(ctx, wordsBatch)
				if err != nil {
					fmt.Sprintf("erreur lors de l'insertion du batch : %v", err)
					return
				}

				compteur += len(wordsBatch)
				fmt.Printf("Processed %d words\n", compteur)

				// Réinitialiser le batch
				wordsBatch = nil
			}
		}
	}
	// Insérer les mots restants
	if len(wordsBatch) > 0 {
		_, err := collection.InsertMany(ctx, wordsBatch)
		if err != nil {
			fmt.Sprintf("erreur lors de l'insertion du batch : %v", err)
			return
		}

		compteur += len(wordsBatch)
		fmt.Printf("Processed %d words\n", compteur)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d mots ont été ajoutés à la base de données.\n", compteur)
}

// Fonctions de nettoyage et validation

func cleanWord(word string) string {
	word = strings.ToLower(word)
	return strings.TrimSpace(word)
}

func isValidWord(word string) bool {
	for _, r := range word {
		if !unicode.IsLetter(r) {
			fmt.Printf("Le mot %s n'est pas valide\n", word)
			return false
		}
	}

	return true
}