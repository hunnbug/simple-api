package main

import (
	"database/sql"
	"main/albumsType"
	"main/batadase"
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
	ctx.IndentedJSON(200, albums)

	albumsType.WriteLogToALogFile(albums, http.Request{Method: "GET"})

}

func postAlbum(ctx *gin.Context) {

	var newAlbum albumsType.Album

	e := ctx.BindJSON(&newAlbum)
	checkErrors(e)

	albumsType.WriteLogToALogFile(newAlbum, http.Request{Method: "POST"})

	batadase.AddToDatabase(db, newAlbum.Name, newAlbum.Year, newAlbum.Artist)

}

func deleteAlbum(_ *gin.Context) {

	deletedAlbum := batadase.DeleteAlbums(db)

	albumsType.WriteLogToALogFile(deletedAlbum, http.Request{Method: "DELETE"})

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
