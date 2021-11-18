package models

import "time"

type base struct {
	ID        uint      `json:"id" gorm:"primarykey,autoincrement"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
