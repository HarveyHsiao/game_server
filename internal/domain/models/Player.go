package models

import "github.com/gorilla/websocket"

type Player struct {
	UserID string `json:"user_id"`
	RoomID string
	Conn   *websocket.Conn
}
