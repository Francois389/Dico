package models

import (
	"sort"
)

type Word struct {
	Word         string `json:"word" bson:"word"`
	Length       int    `json:"length" bson:"length"`
	FirstLetter  string `json:"first_letter" bson:"first_letter"`
	SortedLetter string `json:"sorted_letter" bson:"sorted_letter"`
}

func NewWord(word string) Word {
	return Word{Word: word, Length: len(word), FirstLetter: string(word[0]), SortedLetter: SortLetter(word)}
}

func SortLetter(word string) string {
	var lettres = []rune(word)

	sort.Slice(lettres, func(i, j int) bool {
		return lettres[i] < lettres[j]
	})

	return string(lettres)
}
