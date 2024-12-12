package mot

type Mot struct {
	Word string `json:"word" bson:"word"`
	Length int `json:"length" bson:"length"`
	FirstLetter string `json:"first_letter" bson:"first_letter"`
}
