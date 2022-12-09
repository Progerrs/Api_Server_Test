package store

import (
	"awesomeProject/internal/app/model"
	"database/sql"
)

// UserRepository ...
type UserRepository interface {
	Create(*model.User) error
	Get() []*model.User
	Delete(int) *sql.Row
	Update(Id int, u *model.User) *sql.Row
}
