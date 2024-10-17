package batadase

import "database/sql"

func handleError(e error) {
	if e != nil {
		panic(e)
	}
}

type Album struct {
	Name   string
	Year   string
	Artist string
}

type parsedAlbum struct {
	ID     int
	Name   string
	Year   string
	Artist string
}

func ConnectToDB() *sql.DB {

	db, e := sql.Open("postgres", "host=localhost port=5432 user=postgres password=1 dbname=albums sslmode=disable")
	handleError(e)

	return db

}

func GetAlbums(db *sql.DB) []Album {

	items, e := db.Query("SELECT * FROM albumstable;")
	handleError(e)

	var albums []Album

	for items.Next() {

		var newParsedAlbum parsedAlbum

		e := items.Scan(&newParsedAlbum.ID, &newParsedAlbum.Name, &newParsedAlbum.Year, &newParsedAlbum.Artist)
		handleError(e)

		var newAlbum = Album{
			Name:   newParsedAlbum.Name,
			Year:   newParsedAlbum.Year,
			Artist: newParsedAlbum.Artist,
		}

		albums = append(albums, newAlbum)

	}

	return albums
}

func AddToDatabase(db *sql.DB, name string, year string, artist string) {

	_, e := db.Exec("INSERT INTO albumstable (Name, Year, Artist) VALUES ($1, $2, $3)", name, year, artist)
	handleError(e)

}

func DeleteAlbums(db *sql.DB) Album {

	album, e := db.Query("SELECT * FROM albumstable WHERE ID = (SELECT MAX(ID) FROM albumstable);")
	handleError(e)

	_, e = db.Exec("DELETE FROM albumstable WHERE ID = (SELECT MAX(ID) FROM albumstable);")
	handleError(e)

	var albumToDelete Album

	for album.Next() {
		var newParsedAlbum parsedAlbum

		e := album.Scan(&newParsedAlbum.ID, &newParsedAlbum.Name, &newParsedAlbum.Year, &newParsedAlbum.Artist)
		handleError(e)

		albumToDelete = Album{
			Name:   newParsedAlbum.Name,
			Year:   newParsedAlbum.Year,
			Artist: newParsedAlbum.Artist,
		}
	}

	return albumToDelete

}
