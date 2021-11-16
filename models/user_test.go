package models_test

import (
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
	testcases = []struct {
		name    string
		input   *models.User
		wantErr bool
	}{
		{
			"正しく通ること",
			&models.User{
				ID:        0,
				Name:      "takada",
				Email:     "god@example.com",
				Password:  "kenshi",
				CreatedAt: t,
				UpdatedAt: t,
			},
			false,
		},
		{
			"emailの重複でエラーが発生すること",
			&models.User{
				ID:        0,
				Name:      "mokou",
				Email:     "god@example.com",
				Password:  "yutaka",
				CreatedAt: t,
				UpdatedAt: t,
			},
			true,
		},
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

func TestCreate(t *testing.T) {
	assert := assert.New(t)
	db, teardown := database.SetupTestDB()
	defer teardown()

	var (
		c int64
		i int64 = 1
	)
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := user.Create()
			db.Find(&user).Count(&c)
			if tc.wantErr {
				assert.NotEqual(i, c)
				assert.Error(err)
			} else {
				assert.Equal(i, c)
				assert.NoError(err)
			}
		})
	}
}
