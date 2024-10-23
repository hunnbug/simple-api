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

func WriteLogToALogFile(v any, request http.Request) {

	switch value := v.(type) {

	case Album:

		logBody := fmt.Sprintf("\nTime: %s\nName: %s\nArtist: %s\nYear: %s\nMethod: %s\n--------------------------------------", time.Now().GoString(), value.Name, value.Artist, value.Year, request.Method)

		file, e := os.OpenFile("logs.log", os.O_APPEND|os.O_WRONLY, 0644)
		checkErrors(e)
		defer file.Close()

		file.WriteString(logBody)

	case []Album:

		for _, album := range value {

			logBody := fmt.Sprintf("\nTime: %s\nName: %s\nArtist: %s\nYear: %s", time.Now().GoString(), album.Name, album.Artist, album.Year)

			file, e := os.OpenFile("logs.log", os.O_APPEND|os.O_WRONLY, 0644)
			checkErrors(e)

			file.WriteString(logBody)

		}

		file, e := os.OpenFile("logs.log", os.O_APPEND|os.O_WRONLY, 0644)
		checkErrors(e)
		defer file.Close()

		file.WriteString("\n" + request.Method)
		file.WriteString("\n----------------------------------------------------\n")

	}
}
