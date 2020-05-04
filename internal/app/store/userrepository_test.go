package store_test

import (
	"testing"

	"github.com/reed665/http-rest-api/internal/app/model"
	"github.com/reed665/http-rest-api/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, config.DatabaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, config.DatabaseURL)
	defer teardown("users")

	u := model.TestUser(t)
	email := u.Email

	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
