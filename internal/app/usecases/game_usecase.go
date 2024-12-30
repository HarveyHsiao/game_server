package usecases

import "game_server/internal/domain/repositories"

type GameUsecase interface {
	AddPlayerToRoom(roomID, playerID string) error
	RemovePlayerFromRoom(roomID, playerID string) error
	BroadcaseMessage(roomID, message string) error
}

type gameusecase struct {
	roomRepo repositories.RoomRepository
}

func NewGameUsecase(roomRepo repositories.RoomRepository) GameUsecase {
	return &gameusecase{
		roomRepo: roomRepo,
	}
}

func (u *gameusecase) AddPlayerToRoom(roomID, playerID string) error {
	return nil
}

func (u *gameusecase) RemovePlayerFromRoom(roomID, playerID string) error {
	return nil
}

func (u *gameusecase) BroadcaseMessage(roomID, message string) error {
	return nil
}
