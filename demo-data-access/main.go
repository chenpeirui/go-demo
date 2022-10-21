package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	// Init DB config.
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWD"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	// Open DB.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// Test DB connection.
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	// Queries for album.
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	// Queries for the album.
	newAlb, err := albumById(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", newAlb)

	// Create new album.
	albId, err := addAlbum(Album{
		Title:  "3",
		Artist: "3",
		Price:  3.99,
	})
	if err != nil {
		log.Fatalf("addAlbum %v", err)
	}
	fmt.Printf("ID of added album: %v\n", albId)
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
	var albums []Album

	// Queries rows from db.
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	// Copy rows to []albums.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

// albumById queries for the album with the specified id.
func albumById(id int64) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumById %d: no such id", id)
		}
		return alb, fmt.Errorf("albumById %d: %v", id, err)
	}

	return alb, nil
}

// addAlbum adds the specified album to database,
// returning the album id of the new entry
func addAlbum(alb Album) (int64, error) {
	rst, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	id, err := rst.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}

	return id, nil
}
