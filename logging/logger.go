package logging

import (
	"fmt"
	"main/batadase"
	"net/http"
	"os"
	"time"
)

func WriteLogToALogFile(album batadase.Album, request http.Request) {

	logBody := fmt.Sprintf("time: %s\nName: %s\nArtist: %s\nYear: %s\nMethod: %s\n----------------------------------", time.Now().GoString(), album.Name, album.Artist, album.Year, request.Method)

	e := os.WriteFile("logs.log", []byte(logBody), 0644)
	if e != nil {
		panic(e)
	}

}
