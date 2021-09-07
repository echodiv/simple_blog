package store_test

import (
	"testing"

	"github.com/echodiv/simple_blog/internal/app/models"
	"github.com/echodiv/simple_blog/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryCreate(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(&models.User{
		Email: "user@example.org",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepositoryFindByEmailWithoutUser(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	var email string = "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)
}

func TestUserRepositoryFindByEmailWithtExistingUser(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	var email string = "user@example.org"
	s.User().Create(&models.User{
		Email: email,
	})
	u, err := s.User().FindByEmail(email)
	assert.NotNil(t, u)
	assert.NoError(t, err)
}
