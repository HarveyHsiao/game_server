package main

import "game_server/internal/app/router"

func main() {
	r := router.InitRouter()
	r.Run(":8080")
}
