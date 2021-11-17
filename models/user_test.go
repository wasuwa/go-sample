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

	u := testcases[0].input
	db.Create(u)
	uu, err := u.All()

	assert.Contains(uu, *u)
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
			u := tc.input
			err := u.Create()
			db.Find(u).Count(&c)
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

func TestFind(t *testing.T) {
	assert := assert.New(t)
	db, teardown := database.SetupTestDB()
	defer teardown()

	u1 := testcases[0].input
	db.Create(u1)
	id := int(u1.ID)

	var u2 models.User
	err := u2.Find(id)

	assert.Equal(id, int(u2.ID))
	assert.NoError(err)
}
