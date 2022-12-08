package model

// User ...
type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
	BirthDay  string `json:"birth-day"`
	Gender    string `json:"gender"`
}
