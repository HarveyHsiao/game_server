package usecases

import (
	"fmt"
	"game_server/internal/domain/models"
	"sync"
	"time"
)

type RoomUsecase interface {
	CreateRoom() *models.Room
	JoinRoom(roomID string, player *models.Player) error
	LeftRoom(player *models.Player) error
	GetRoom(roomID string) (*models.Room, error)
}

type roomUsecase struct {
	rooms map[string]*models.Room
	mu    sync.Mutex
}

func NewRoomUsecase() RoomUsecase {
	return &roomUsecase{
		rooms: make(map[string]*models.Room),
	}
}

func (u *roomUsecase) CreateRoom() *models.Room {
	u.mu.Lock()
	defer u.mu.Unlock()

	roomID := generateRoomID()
	room := &models.Room{
		ID:      generateRoomID(),
		Players: make(map[string]*models.Player),
	}

	u.rooms[roomID] = room

	return room
}

func (u *roomUsecase) JoinRoom(roomID string, player *models.Player) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	room, exist := u.rooms[roomID]

	if !exist {
		return fmt.Errorf("room not found")
	}

	if len(room.Players) >= 4 {
		return fmt.Errorf("room is full")
	}

	player.InsideRoom = true
	player.RoomID = roomID
	room.Players[player.UserID] = player

	return nil
}

func (u *roomUsecase) LeftRoom(player *models.Player) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if !player.InsideRoom {
		return fmt.Errorf("player is not inside room")
	}

	room, exist := u.rooms[player.RoomID]

	if !exist {
		return fmt.Errorf("room not found")
	}

	delete(room.Players, player.UserID)
	player.InsideRoom = false

	if len(room.Players) == 0 {
		delete(u.rooms, room.ID)
		fmt.Printf("close room ID: %s", room.ID)
	}

	return nil
}

func (u *roomUsecase) GetRoom(roomID string) (*models.Room, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	room, exist := u.rooms[roomID]

	if !exist {
		return nil, fmt.Errorf("room ID: %s doesn't existed", roomID)
	}

	return room, nil
}

// random RoomID
func generateRoomID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
