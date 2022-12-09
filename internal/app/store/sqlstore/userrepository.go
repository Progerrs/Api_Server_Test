package sqlstore

import (
	"awesomeProject/internal/app/model"
	"database/sql"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// Create ...
func (r *UserRepository) Create(u *model.User) error {
	return r.store.db.QueryRow(
		"INSERT INTO public.users(firstname, lastname, birthday, gender) "+
			"VALUES ($1, $2, $3, $4) RETURNING Id",
		u.FirstName,
		u.LastName,
		u.BirthDay,
		u.Gender,
	).Scan(&u.Id)
}

// Get ...
func (r *UserRepository) Get() *sql.Row {
	return r.store.db.QueryRow("SELECT row_to_json(users) FROM users")
}

// Delete ...
func (r *UserRepository) Delete(Id int) *sql.Row {
	return r.store.db.QueryRow("DELETE FROM users WHERE id = $1",
		Id)
}

// Update ...
func (r *UserRepository) Update(Id int, u *model.User) *sql.Row {
	return r.store.db.QueryRow("UPDATE users SET firstname = $2, lastname = $3, birthday = $4, gender = $5"+
		"WHERE id = $1",
		Id,
		u.FirstName,
		u.LastName,
		u.BirthDay,
		u.Gender)
}
