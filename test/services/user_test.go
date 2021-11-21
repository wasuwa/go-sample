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
	t    time.Time
	base = &models.Base{
		ID:        0,
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
		input   *models.ReceiveUser
		wantErr bool
	}{
		{
			"正しく通ること",
			&models.ReceiveUser{
				Name:     "mokou",
				Email:    "baba@example.com",
				Password: "yutaka",
			},
			false,
		},
		{
			"emailの重複でエラーが発生すること",
			&models.ReceiveUser{
				Name:     "mokou",
				Email:    "baba@example.com",
				Password: "yutaka",
			},
			true,
		},
	}
)

func init() {
	config.ResetPath()
}

func TestAllUser(t *testing.T) {
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

func TestFindUser(t *testing.T) {
	assert := assert.New(t)
	db, teardown := database.SetupTestDB()
	defer teardown()

	u, err := services.FindUser(0)
	assert.Nil(u)
	assert.Error(err)

	db.Create(user)
	id := int(user.ID)
	u, err = services.FindUser(id)
	assert.Equal(user, u)
	assert.NoError(err)
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	_, teardown := database.SetupTestDB()
	defer teardown()

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			u, err := services.CreateUser(tc.input)
			if tc.wantErr {
				assert.Nil(u)
				assert.Error(err)
			} else {
				assert.Equal(tc.input.Name, u.Name)
				assert.NoError(err)
			}
		})
	}
}

// func TestUpdate(t *testing.T) {
// 	assert := assert.New(t)
// 	db, teardown := database.SetupTestDB()
// 	defer teardown()

// 	db.Create(user)
// 	user.ID    = 0
// 	user.Email = "god@example.com"
// 	db.Create(user)
// 	id := int(user.ID)

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			err := services.UpdateUser(tc.input, id)
// 			if tc.wantErr {
// 				assert.Error(err)
// 			} else {
// 				assert.NoError(err)
// 			}
// 		})
// 	}
// }
