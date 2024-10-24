package main

import (
	"main/batadase"
	server "main/server"

	_ "github.com/lib/pq"
)

func main() {

	db := batadase.ConnectToDB()

	server.StartServer(db)

}
