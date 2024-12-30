package repositories

import "game_server/internal/domain/models"

type RoomRepository interface {
	CreateRoom(room *models.Room) error
	GetRoomByID(id string) (*models.Room, error)
	SaveRoom(room *models.Room) error
}
