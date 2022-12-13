package sqlstore

import (
	"awesomeProject/internal/app/model"
	"database/sql"
	"log"
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
func (r *UserRepository) Get() []*model.User {
	rows, err := r.store.db.Query("SELECT id, firstname, lastname, birthday, gender " +
		"FROM users")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return nil
	}
	defer rows.Close()
	users := make([]*model.User, 0)
	for rows.Next() {
		u := new(model.User)
		err := rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.BirthDay, &u.Gender)
		if err != nil {
			log.Fatal(err)
		}
		errConvert := u.BeforeGet()
		if errConvert != nil {
			log.Fatal(errConvert)
		}
		users = append(users, u)
	}
	return users
}

// Delete ...
func (r *UserRepository) Delete(Id int) *sql.Row {
	return r.store.db.QueryRow("DELETE FROM users WHERE id = $1",
		Id)
}

// Update ...
func (r *UserRepository) Update(Id string, u *model.User) error {
	return r.store.db.QueryRow("UPDATE users SET firstname = $2, lastname = $3, birthday = $4, gender = $5"+
		"WHERE id = $1  RETURNING Id",
		Id,
		u.FirstName,
		u.LastName,
		u.BirthDay,
		u.Gender).Scan(&u.Id)
}
