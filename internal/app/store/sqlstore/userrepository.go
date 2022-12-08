package sqlstore

import (
	"awesomeProject/internal/app/model"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	return r.store.db.QueryRow(
		"INSERT INTO public.users(firstname, lastname, birthday, gender) VALUES ($1, $2, $3, $4) RETURNING Id",
		u.FirstName,
		u.LastName,
		u.BirthDay,
		u.Gender,
	).Scan(&u.Id)
}
