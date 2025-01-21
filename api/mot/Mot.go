package mot

import (
	"sort"
)

type Mot struct {
	Word         string `json:"word" bson:"word"`
	Length       int    `json:"length" bson:"length"`
	FirstLetter  string `json:"first_letter" bson:"first_letter"`
	SortedLetter string `json:"sorted_letter" bson:"sorted_letter"`
}

func NewMot(word string) Mot {
	return Mot{Word: word, Length: len(word), FirstLetter: string(word[0]), SortedLetter: sortLetter(word)}
}

func sortLetter(word string) string {
	var lettres = []rune(word)

	sort.Slice(lettres, func(i, j int) bool {
		return lettres[i] < lettres[j]
	})

	return string(lettres)
}
