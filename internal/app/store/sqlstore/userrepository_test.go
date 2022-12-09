package sqlstore_test

import (
	"awesomeProject/internal/app/model"
	"awesomeProject/internal/app/store/sqlstore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, _ := sqlstore.TestDB(t, databaseURL)

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u.Id)
}

func TestUserRepository_Get(t *testing.T) {
	db, _ := sqlstore.TestDB(t, databaseURL)

	s := sqlstore.New(db)
	assert.NotNil(t, s.User().Get())
}

func TestUserRepository_Delete(t *testing.T) {
	db, _ := sqlstore.TestDB(t, databaseURL)

	s := sqlstore.New(db)
	assert.NotNil(t, s.User().Delete(12))
}

func TestUserRepository_Update(t *testing.T) {
	db, _ := sqlstore.TestDB(t, databaseURL)

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NotNil(t, s.User().Update(13, u))
}
