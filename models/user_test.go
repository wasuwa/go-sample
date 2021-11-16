package models_test

import (
	"fmt"
	"testing"
	"time"
	"twitter-app/database"
	"twitter-app/models"

	"github.com/stretchr/testify/assert"
)

var (
	t	time.Time
	user = &models.User{
		ID:        0,
		Name:      "takada",
		Email:     "god@example.com",
		Password:  "kenshi",
		CreatedAt: t,
		UpdatedAt: t,
	}
)

func TestAll(t *testing.T) {
	assert := assert.New(t)
	db, teardown := database.SetupTestDB()
	defer teardown()

	db.Create(user)
	uu, err := user.All()

	assert.Contains(uu, *user)
	assert.NoError(err)
}
