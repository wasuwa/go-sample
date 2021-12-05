package models

import "time"

type Base struct {
	ID        uint      `json:"id" gorm:"primarykey,autoincrement"`
	CreatedAt time.Time `json:"createdAt"`
}
