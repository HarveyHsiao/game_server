package models

type Room struct {
	ID      string `json:"room_id"`
	Players map[string]*Player
}
