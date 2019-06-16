package main

import (
	"github.com/rinaldypasya/simple-messaging/api"
	"github.com/rinaldypasya/simple-messaging/config"
)

var (
	port = ":8000"
)

func main() {
	db := config.DBInit()
	inDB := &api.InDB{DB: db}

	server := api.InitRouter(inDB)
	server.Run(port)
}
