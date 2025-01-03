package router

import (
	"game_server/internal/app/handlers"
	"game_server/internal/app/usecases"
	"game_server/internal/domain/repositories"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	// init Gin
	r := gin.Default()

	// repo
	roomRep := repositories.NewRoomRepository(nil, nil)

	// usecase
	roomUsecase := usecases.NewRoomUsecase(roomRep)

	// handler
	roomHandler := handlers.NewRoomHandler(roomUsecase)

	// router
	roomApi := r.Group("room")

	roomApi.POST("/create", func(context *gin.Context) {
		roomHandler.CreateRoom(context)
	})

	//roomApi.POST("/")

	return r
}
