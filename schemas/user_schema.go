package schema

type UserCreation struct {
	Firstname string `json:"firstname"` // Correct: the key is "json", and the value is "firstname"
	Lastname  string `json:"lastname"`
}