package main

import (
	"main/batadase"
	"main/server"

	_ "github.com/lib/pq"
)

func main() {

	batadase.ConnectToDB()

	server.StartServer()

}
