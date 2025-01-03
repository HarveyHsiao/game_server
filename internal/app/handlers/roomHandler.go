package handlers

import (
	"fmt"
	"game_server/internal/app/usecases"
	"game_server/internal/domain/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type roomHandler struct {
	roomUsecase usecases.RoomUsecase
	upgrader    websocket.Upgrader
}

func NewRoomHandler(usecase usecases.RoomUsecase) *roomHandler {
	return &roomHandler{
		roomUsecase: usecase,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			}},
	}
}

func (h *roomHandler) CreateRoom(c *gin.Context) {
	room, err := h.roomUsecase.CreateRoom()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"room": room})
}

func (h *roomHandler) JoinRoom(c *gin.Context) {
	roomID := c.Query("id")

	_, err := h.roomUsecase.GetRoom(roomID)

	if err != nil {
		fmt.Println(roomID)
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	conn, socketErr := h.upgrader.Upgrade(c.Writer, c.Request, nil)

	if socketErr != nil {
		fmt.Println("aa")
		log.Println(socketErr)
		//c.JSON(http.StatusNotFound, gin.H{"error": socketErr})
		return
	}

	playerID := c.Query("player_id")

	player := &models.Player{
		UserID: playerID,
		RoomID: roomID,
		Conn:   conn,
	}

	err = h.roomUsecase.JoinRoom(roomID, player)

	if err != nil {
		fmt.Println("ccc")
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	go func() {
		defer h.roomUsecase.LeftRoom(player)

		for {
			_, msg, connErr := conn.ReadMessage()
			if connErr != nil {
				break
			}
			h.roomUsecase.BroadCast(msg, player)
		}
	}()

}
