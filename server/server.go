package server

import (
	"main/albumsType"
	"main/batadase"
	"main/loggingSystem"

	"github.com/gin-gonic/gin"
)

func checkErrors(e error) {

	if e != nil {
		panic(e)
	}

}

func getAlbums(ctx *gin.Context) {

	albums := batadase.GetAlbums()
	ctx.IndentedJSON(200, albums)

	loggingSystem.WriteLogToALogFile(albums, *ctx.Request)

}

func postAlbum(ctx *gin.Context) {

	var newAlbum albumsType.Album

	e := ctx.BindJSON(&newAlbum)
	checkErrors(e)

	loggingSystem.WriteLogToALogFile(newAlbum, *ctx.Request)

	batadase.AddToDatabase(newAlbum.Name, newAlbum.Year, newAlbum.Artist)

}

func deleteAlbum(ctx *gin.Context) {

	deletedAlbum := batadase.DeleteAlbums()

	loggingSystem.WriteLogToALogFile(deletedAlbum, *ctx.Request)

}

func StartServer() {

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.DELETE("/albums", deleteAlbum)

	router.Run("192.168.1.78:8080")

}
