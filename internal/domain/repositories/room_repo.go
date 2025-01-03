package repositories

import (
	"game_server/internal/domain/models"
	"gorm.io/gorm"
)

type RoomRepository interface {
	CreateRoom(room *models.Room) error
	GetRoomByID(id string) (*models.Room, error)
	SaveRoom(room *models.Room) error
}

type roomRepository struct {
	WriteDB *gorm.DB
	ReadDB  *gorm.DB
}

func NewRoomRepository(WriteDB *gorm.DB, ReadDB *gorm.DB) RoomRepository {
	return &roomRepository{
		WriteDB: WriteDB,
		ReadDB:  ReadDB,
	}
}

func (r *roomRepository) CreateRoom(room *models.Room) error {
	return nil
}

func (r *roomRepository) GetRoomByID(id string) (*models.Room, error) {
	return nil, nil
}

func (r *roomRepository) SaveRoom(room *models.Room) error {
	return nil
}
