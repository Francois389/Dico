package mot

type Mot struct {
	Word string `json:"word" bson:"word"`
	Length int `json:"length" bson:"length"`
	FirstLetter string `json:"first_letter" bson:"first_letter"`
}

func NewMot(word string) Mot {
	return Mot{Word: word, Length: len(word), FirstLetter: string(word[0]), SortedLetter: sortLetter(word)}
}