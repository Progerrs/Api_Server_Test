package model

import "testing"

// TestUser ...
func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		FirstName: "Test",
		LastName:  "Test",
		BirthDay:  "8-12-2022",
		Gender:    "Male",
	}
}
