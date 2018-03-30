package models

// Person represents a person node in the family tree with all
// the necessary information like name, birthdate and so on.
type Person struct {
	Key     string `json:"key"`
	Payload struct {
		FirstName    string `json:"firstName"`
		LastName     string `json:"lastName"`
		MaidenName   string `json:"maidenName"`
		Birthday     string `json:"birthday"`
		Death        string `json:"death"`
		PlaceOfBirth string `json:"placeOfBirth"`
		Notes        string `json:"notes"`
	} `json:"payload"`
}
