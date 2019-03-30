package entity

type Questions struct {
	Question string   `json:"question"`
	Correct  string   `json:"correct"`
	Answers  []string `json:"answers"`
}

type UserRanking struct {
	Name           string
	Percentage     string
	CorrectAnswers float64
}
