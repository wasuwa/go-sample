package models_test

import (
	"testing"
	"time"
	"twitter-app/config"
	"twitter-app/database"
	"twitter-app/models"
	"twitter-app/services"

	"github.com/stretchr/testify/assert"
)

var (
	t time.Time
	base = &models.Base {
		ID: 0,
		CreatedAt: t,
		UpdatedAt: t,
	}
	user = &models.User{
		Name:     "takada",
		Email:    "god@example.com",
		Password: "kenshi",
	}
	testcases = []struct {
		name    string
		input   *models.User
		wantErr bool
	}{
		{
			"正しく通ること",
			&models.User{
				Name:      "takada",
				Email:     "god@example.com",
				Password:  "kenshi",
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

func init() {
	config.ResetPath()
}

func TestAll(t *testing.T) {
	assert := assert.New(t)
	db, teardown := database.SetupTestDB()
	defer teardown()

	users, err := services.AllUser()
	assert.Nil(users)
	assert.Error(err)

	db.Create(user)
	users, err = services.AllUser()
	assert.Contains(users, *user)
	assert.NoError(err)
}

// func TestCreate(t *testing.T) {
// 	assert := assert.New(t)
// 	db, teardown := database.SetupTestDB()
// 	defer teardown()

// 	var (
// 		c int64
// 		i int64 = 1
// 	)
// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			u := tc.input
// 			err := u.Create()
// 			db.Find(u).Count(&c)
// 			if tc.wantErr {
// 				assert.NotEqual(i, c)
// 				assert.Error(err)
// 			} else {
// 				assert.Equal(i, c)
// 				assert.NoError(err)
// 			}
// 		})
// 	}
// }

// func TestFind(t *testing.T) {
// 	assert := assert.New(t)
// 	db, teardown := database.SetupTestDB()
// 	defer teardown()

// 	db.Create(user)
// 	id := int(user.ID)

// 	var (
// 		u2 models.User
// 		u3 models.User
// 	)

// 	err := u2.Find(id)
// 	assert.Equal(id, int(u2.ID))
// 	assert.NoError(err)

// 	err = u3.Find(0)
// 	assert.Error(err)
// }

// func TestUpdate(t *testing.T) {
// 	assert := assert.New(t)
// 	db, teardown := database.SetupTestDB()
// 	defer teardown()

// 	db.Create(user)
// 	id := int(user.ID)

// 	user.Email = "mokou@example.com"
// 	db.Create(user)

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			err := tc.input.Update(id)
// 			if tc.wantErr {
// 				assert.Error(err)
// 			} else {
// 				assert.NoError(err)
// 			}
// 		})
// 	}
// }
