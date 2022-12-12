package model

import (
	"strings"
	"time"
)

// User ...
type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDay  string `json:"birth_day"`
	Gender    string `json:"gender"`
	Years     int    `json:"years"`
}

func (u *User) BeforeGet() error {
	u.BirthDay = strings.TrimRight(u.BirthDay, "T00:00:00Z")
	now := time.Now()
	birthday, err := time.Parse("2006-01-2", u.BirthDay)
	if err != nil {
		return err
	}
	u.Years = int(now.Sub(birthday).Hours() / 24 / 365)
	return nil
}
