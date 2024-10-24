package server

import (
	"database/sql"
	"main/albumsType"
	"main/batadase"
	"main/loggingSystem"

	"github.com/gin-gonic/gin"
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

	loggingSystem.WriteLogToALogFile(albums, *ctx.Request)

}

func postAlbum(ctx *gin.Context) {

	var newAlbum albumsType.Album

	e := ctx.BindJSON(&newAlbum)
	checkErrors(e)

	loggingSystem.WriteLogToALogFile(newAlbum, *ctx.Request)

	batadase.AddToDatabase(db, newAlbum.Name, newAlbum.Year, newAlbum.Artist)

}

func deleteAlbum(ctx *gin.Context) {

	deletedAlbum := batadase.DeleteAlbums(db)

	loggingSystem.WriteLogToALogFile(deletedAlbum, *ctx.Request)

}

func StartServer(a *sql.DB) {

	db = a

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.DELETE("/albums", deleteAlbum)

	router.Run("localhost:8080")

}
