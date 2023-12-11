package dataBase

import (
	"log"
)

func CreateTable() error {
	query := `DROP TABLE IF EXISTS album;
	CREATE TABLE album (
		id SERIAL PRIMARY KEY,
		album_title VARCHAR(255) NOT NULL,
		artist_name VARCHAR(255) NOT NULL
	);`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("failed to create album table: %v", err)
	}
	return err
}

func GetRecords() {
	query := `SELECT * FROM album;`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to get Records: %v", err)
	}

}

func InputRecords(title, artist string) error {
	album := NewAlbum(title, artist)
	query := `INSERT INTO album (album_title, artist_name) VALUES ($1, $2);`
	_, err := DB.Exec(query, album.albumTitle, album.artistName)
	return err
}
