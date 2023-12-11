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

// C
func InputRecords(title, artist string) error {
	album := NewAlbum(title, artist)
	query := `INSERT INTO album (album_title, artist_name) VALUES ($1, $2);`
	_, err := DB.Exec(query, album.albumTitle, album.artistName)
	if err != nil {
		log.Printf("Error Creating new record: %v", err)
		return err
	}
	return err
}

// R
func GetRecords() {
	query := `SELECT * FROM album;`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to get Records: %v", err)
	}
}

// U
func UpdateRecords(newAlbum, artistName string) error {
	query := `UPDATE album SET album_title = $1 WHERE artist_name = $2;`
	_, err := DB.Exec(query, newAlbum, artistName)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// D
func RemoveRecords(artistName string) error {
	query := `DELETE FROM album WHERE artist_name = $1;`
	_, err := DB.Exec(query, artistName)
	if err != nil {
		log.Printf("Unable to remove record: %v", err)
		return err
	}
	return nil
}
