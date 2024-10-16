package main

import (
	"database/sql"
	"main/batadase"

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
}

func postAlbum(ctx *gin.Context) {

	var newAlbum batadase.Album

	e := ctx.BindJSON(&newAlbum)
	checkErrors(e)
}

func deleteAlbum(_ *gin.Context) {

	batadase.DeleteAlbums(db)

}

func startServer() {

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.DELETE("/albums", deleteAlbum)

	router.Run("192.168.1.78:8080")

}

func main() {

	db = batadase.ConnectToDB()
	startServer()

}
