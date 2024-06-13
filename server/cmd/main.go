package main

import (
	"github.com/kviatkovsky/ChatWS/db"
	"github.com/kviatkovsky/ChatWS/internal/user"
	"github.com/kviatkovsky/ChatWS/internal/ws"
	"github.com/kviatkovsky/ChatWS/router"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		panic(err)
	}

	userRep := user.NewRepository(db.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8080")
}
