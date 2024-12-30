package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	UserID     string `json:"user_id"`
	InsideRoom bool
	RoomID     string
}
