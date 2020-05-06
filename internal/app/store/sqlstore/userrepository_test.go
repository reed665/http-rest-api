package sqlstore_test

import (
	"testing"

	"github.com/reed665/http-rest-api/internal/app/model"
	"github.com/reed665/http-rest-api/internal/app/store"
	"github.com/reed665/http-rest-api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, conf.DatabaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, conf.DatabaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	email := u.Email

	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
