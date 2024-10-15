package main

import (
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Year   string `json:"year"`
	Artist string `json:"artist"`
}

var albums = []album{
	{ID: 1, Name: "AM", Year: "2013", Artist: "Arctic Monkeys"},
	{ID: 2, Name: "A Place For Us To Dream", Year: "2016", Artist: "Placebo"},
}

func checkErrors(e error) {

	if e != nil {
		panic(e)
	}

	os.WriteFile("l.log", []byte("an error occured! "+e.Error()), 0644)

}

func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(200, albums)
}

func postAlbum(ctx *gin.Context) {

	var newAlbum album

	e := ctx.BindJSON(&newAlbum)
	checkErrors(e)

	var lastAlbum album

	if len(albums) != 0 {
		lastAlbum = albums[len(albums)-1]
	} else {
		lastAlbum.ID = 0
	}

	newAlbum.ID = lastAlbum.ID + 1

	albums = append(albums, newAlbum)
	ctx.IndentedJSON(201, newAlbum)

}

func deleteAlbum(_ *gin.Context) {

	switch {

	case len(albums) >= 2:
		albums = albums[:len(albums)-2]

	case len(albums) == 1:
		albums = nil

	case len(albums) == 0:
		return

	}

}

func startServer() {

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.DELETE("/albums", deleteAlbum)

	router.Run("192.168.1.78:8080")

}

func connectToSQL() {

	db, e := sql.Open("postgres", "host=localhost port=5432 user=postgres password=1 dbname=test sslmode=disable")
	checkErrors(e)

	db.Ping()

}

func main() {

	startServer()

}
