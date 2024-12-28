package models

type Room struct {
	RoomID  string `json:"room_id"`
	Players map[string]*Player
}
