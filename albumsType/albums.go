package albumsType

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type Album struct {
	Name   string
	Year   string
	Artist string
}

type ParsedAlbum struct {
	ID     int
	Name   string
	Year   string
	Artist string
}

func checkErrors(e error) {
	if e != nil {
		panic(e)
	}
}

type Log interface {
	WriteLogToALogFile(http.Request)
}

func (album Album) WriteLogToALogFile(request http.Request) {

	logBody := fmt.Sprintf("\nTime: %s\nName: %s\nArtist: %s\nYear: %s\nMethod: %s\n--------------------------------------", time.Now().GoString(), album.Name, album.Artist, album.Year, request.Method)

	file, e := os.OpenFile("logs.log", os.O_APPEND|os.O_WRONLY, 0644)
	checkErrors(e)
	defer file.Close()

	file.WriteString(logBody)

}

type albums []Album

func (al albums) WriteLogToALogFile(request http.Request) {

	for _, album := range al {

		logBody := fmt.Sprintf("\nTime: %s\nName: %s\nArtist: %s\nYear: %s\nMethod: %s", time.Now().GoString(), album.Name, album.Artist, album.Year, request.Method)

		file, e := os.OpenFile("logs.log", os.O_APPEND|os.O_WRONLY, 0644)
		checkErrors(e)

		file.WriteString(logBody)

	}

	file, e := os.OpenFile("logs.log", os.O_APPEND|os.O_WRONLY, 0644)
	checkErrors(e)
	defer file.Close()

	file.WriteString("----------------------------------------------------")

}
