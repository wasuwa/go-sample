package models_test

import (
	"testing"
	"time"
	"twitter-app/database"
	"twitter-app/models"

	"github.com/stretchr/testify/assert"
)

var (
	user = &models.User{
		ID:        0,
		Name:      "jun",
		Email:     "katou@example.com",
		Password:  "nejiki",
		CreatedAt: t,
		UpdatedAt: t,
	}
	t         time.Time
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
		// {
		// 	"emailの重複でエラーが発生すること",
		// 	&models.User{
		// 		ID:        0,
		// 		Name:      "mokou",
		// 		Email:     "god@example.com",
		// 		Password:  "yutaka",
		// 		CreatedAt: t,
		// 		UpdatedAt: t,
		// 	},
		// 	true,
		// },
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

	db.Create(user)
	id := int(user.ID)

	var (
		u2 models.User
		u3 models.User
	)

	err := u2.Find(id)
	assert.Equal(id, int(u2.ID))
	assert.NoError(err)

	err = u3.Find(0)
	assert.Error(err)
}

func TestUpdate(t *testing.T) {
	assert := assert.New(t)
	db, teardown := database.SetupTestDB()
	defer teardown()

	db.Create(user)
	id := int(user.ID)

	user.Email = "mokou@example.com"
	db.Create(user)

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.Update(id)
			if tc.wantErr {
				assert.Error(err)
			} else {
				assert.NoError(err)
			}
		})
	}
}
