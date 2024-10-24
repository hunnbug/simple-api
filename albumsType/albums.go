package albumsType

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
