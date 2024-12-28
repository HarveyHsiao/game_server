package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	userID string `json:"user_id"`
}
