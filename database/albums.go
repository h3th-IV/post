package database

type Album struct {
	albumTitle string
	artistName string
}

// func InputRecords() error {
// 	query := `INSERT INTO album (album_title, artist_name) VALUES ('King''s Disease', 'Nas');`
// 	_, err := DB.Exec(query)
// 	return err
// }

func NewAlbum(title, artist string) *Album {
	New := Album{
		albumTitle: title,
		artistName: artist,
	}
	return &New
}
