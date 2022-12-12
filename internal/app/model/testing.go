package model

import "testing"

// TestUser ...
func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		FirstName: "First",
		LastName:  "Test",
		BirthDay:  "6-25-2012",
		Gender:    "Male",
	}
}
