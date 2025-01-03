package usecases

import (
	"fmt"
	"game_server/internal/domain/models"
	"game_server/internal/domain/repositories"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type RoomUsecase interface {
	CreateRoom() (*models.Room, error)
	JoinRoom(roomID string, player *models.Player) error
	LeftRoom(player *models.Player) error
	GetRoom(roomID string) (*models.Room, error)
	BroadCast(msg []byte, player *models.Player) error
}

type roomUsecase struct {
	roomRepo repositories.RoomRepository
	rooms    map[string]*models.Room
	mu       sync.Mutex
}

func NewRoomUsecase(roomRepo repositories.RoomRepository) RoomUsecase {
	return &roomUsecase{
		roomRepo: roomRepo,
		rooms:    make(map[string]*models.Room),
	}
}

func (u *roomUsecase) CreateRoom() (*models.Room, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	roomID := generateRoomID()

	room, exist := u.rooms[roomID]

	if exist {
		return room, fmt.Errorf("room has existed")
	}

	newRoom := &models.Room{
		ID:      generateRoomID(),
		Players: make(map[string]*models.Player),
	}

	u.rooms[roomID] = room

	return newRoom, nil
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

	player.RoomID = roomID
	room.Players[player.UserID] = player

	return nil
}

func (u *roomUsecase) LeftRoom(player *models.Player) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	room, exist := u.rooms[player.RoomID]

	if !exist {
		return fmt.Errorf("room not found")
	}

	delete(room.Players, player.UserID)

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

func (u *roomUsecase) BroadCast(msg []byte, player *models.Player) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	room, exist := u.rooms[player.RoomID]

	if !exist {
		return fmt.Errorf("room ID: %s doesn't existed", room.ID)
	}

	for _, roomPlayer := range room.Players {
		roomPlayer.Conn.WriteMessage(websocket.TextMessage, msg)
	}

	return nil
}

// random RoomID
func generateRoomID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

