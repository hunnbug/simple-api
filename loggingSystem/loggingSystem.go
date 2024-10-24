package loggingSystem

import (
	"fmt"
	"main/albumsType"
	"net/http"
	"os"
	"time"
)

func checkErrors(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteLogToALogFile(v any, request http.Request) {

	switch value := v.(type) {

	case albumsType.Album:

		logBody := fmt.Sprintf("\nName: %s\nArtist: %s\nYear: %s\n\nTime: %s\nRequest: %s\n--------------------------------------\n", value.Name, value.Artist, value.Year, time.Now().GoString(), request.Method)

		file, e := os.OpenFile("logs.log", os.O_APPEND|os.O_WRONLY, 0644)
		checkErrors(e)
		defer file.Close()

		file.WriteString(logBody)

	case []albumsType.Album:

		for _, album := range value {

			logBody := fmt.Sprintf("\nName: %s\nArtist: %s\nYear: %s\n", album.Name, album.Artist, album.Year)

			file, e := os.OpenFile("logs.log", os.O_APPEND|os.O_WRONLY, 0644)
			checkErrors(e)

			file.WriteString(logBody)

		}

		file, e := os.OpenFile("logs.log", os.O_APPEND|os.O_WRONLY, 0644)
		checkErrors(e)
		defer file.Close()

		file.WriteString("\nTime: " + time.Now().GoString())
		file.WriteString("\nrequest: " + request.Method)
		file.WriteString("\n----------------------------------------------------\n")

	}
}
