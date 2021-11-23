package services_test

import (
	"testing"
	"twitter-app/config"
	"twitter-app/database"
	"twitter-app/models"
	"twitter-app/services"

	"github.com/stretchr/testify/assert"
)

var (
	user = &models.User{
		Name:     "takada",
		Email:    "god@example.com",
		Password: "kenshi",
	}
	tests = []struct {
		name    string
		input   *models.ReceiveUser
		wantErr bool
	}{
		{
			"正しく通ること",
			&models.ReceiveUser{
				Name:     "takada",
				Email:    "god@example.com",
				Password: "kenshi",
			},
			false,
		},
		{
			"Emailが存在しないとエラーが返ること",
			&models.ReceiveUser{
				Name:     "takada",
				Email:    "gomi@example.com",
				Password: "kenshi",
			},
			true,
		},
	}
)

func init() {
	config.ResetPath()
}

func TestSearchUser(t *testing.T) {
	assert := assert.New(t)
	db, teardown := database.SetupTestDB()
	defer teardown()

	db.Create(user)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := services.SearchUser(tt.input)
			if tt.wantErr {
				assert.Error(err)
			} else {
				assert.NoError(err)
				assert.Equal(user, u)
			}
		})
	}
}
