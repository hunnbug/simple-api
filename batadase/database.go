package batadase

import (
	"database/sql"
	"main/albumsType"
)

var DB *sql.DB

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

func ConnectToDB() {

	var e error

	DB, e = sql.Open("postgres", "host=localhost port=5432 user=postgres password=1 dbname=albums sslmode=disable")
	handleError(e)

	DB.Ping()
	handleError(e)

}

func GetAlbums() []albumsType.Album {

	items, e := DB.Query("SELECT * FROM albumstable;")
	handleError(e)

	var albums []albumsType.Album

	for items.Next() {

		var newParsedAlbum albumsType.ParsedAlbum

		e := items.Scan(&newParsedAlbum.ID, &newParsedAlbum.Name, &newParsedAlbum.Year, &newParsedAlbum.Artist)
		handleError(e)

		var newAlbum = albumsType.Album{
			Name:   newParsedAlbum.Name,
			Year:   newParsedAlbum.Year,
			Artist: newParsedAlbum.Artist,
		}

		albums = append(albums, newAlbum)

	}

	return albums
}

func AddToDatabase(name string, year string, artist string) {

	_, e := DB.Exec("INSERT INTO albumstable (Name, Year, Artist) VALUES ($1, $2, $3)", name, year, artist)
	handleError(e)

}

func DeleteAlbums() albumsType.Album {

	album, e := DB.Query("SELECT * FROM albumstable WHERE ID = (SELECT MAX(ID) FROM albumstable);")
	handleError(e)

	_, e = DB.Exec("DELETE FROM albumstable WHERE ID = (SELECT MAX(ID) FROM albumstable);")
	handleError(e)

	var albumToDelete albumsType.Album

	for album.Next() {
		var newParsedAlbum albumsType.ParsedAlbum

		e := album.Scan(&newParsedAlbum.ID, &newParsedAlbum.Name, &newParsedAlbum.Year, &newParsedAlbum.Artist)
		handleError(e)

		albumToDelete = albumsType.Album{
			Name:   newParsedAlbum.Name,
			Year:   newParsedAlbum.Year,
			Artist: newParsedAlbum.Artist,
		}
	}

	return albumToDelete

}
