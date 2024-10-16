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

		type parsedAlbum struct {
			ID     int
			Name   string
			Year   string
			Artist string
		}

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

func PostAlbums(db *sql.DB) {

}

func DeleteAlbums(db *sql.DB) {

	_, e := db.Exec("DELETE FROM albumstable WHERE ID_Album = (SELECT MAX(ID_Album) FROM albumstable);")
	handleError(e)

}
