package store

import "awesomeProject/internal/app/model"

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
}
