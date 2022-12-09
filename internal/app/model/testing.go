package model

import "testing"

// TestUser ...
func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		FirstName: "Test",
		LastName:  "Test",
		BirthDay:  "9-12-2022",
		Gender:    "Male",
	}
}
