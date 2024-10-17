package main

import (
	"database/sql"
	"main/batadase"
	"main/logging"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

var db *sql.DB

func checkErrors(e error) {

	if e != nil {
		panic(e)
	}

}

func getAlbums(ctx *gin.Context) {

	albums := batadase.GetAlbums(db)
	for _, album := range albums {
		logging.WriteLogToALogFile(album, http.Request{Method: "GET"})
	}

	ctx.IndentedJSON(200, albums)

}

func postAlbum(ctx *gin.Context) {

	var newAlbum batadase.Album

	e := ctx.BindJSON(&newAlbum)
	checkErrors(e)

	batadase.AddToDatabase(db, newAlbum.Name, newAlbum.Year, newAlbum.Artist)

	logging.WriteLogToALogFile(newAlbum, http.Request{Method: "POST"})

}

func deleteAlbum(_ *gin.Context) {

	deletedAlbum := batadase.DeleteAlbums(db)
	logging.WriteLogToALogFile(deletedAlbum, http.Request{Method: "DELETE"})

}

func startServer() {

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.DELETE("/albums", deleteAlbum)

	router.Run("localhost:8080")

}

func main() {

	db = batadase.ConnectToDB()

	startServer()

}
