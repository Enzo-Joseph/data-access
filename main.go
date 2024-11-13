package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {
	db, err := gorm.Open(sqlite.Open("recordings.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Album{})
	// db.Create(&Album{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99})
	// db.Create(&Album{ID: 2, Title: "Giant Steps", Artist: "John Coltrane", Price: 63.99})
	// db.Create(&Album{ID: 3, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99})
	// db.Create(&Album{ID: 4, Title: "Sarah Vaughan", Artist: "Sarah Vaughan", Price: 34.98})

	albums, err := albumsByArtist(db, "John Coltrane")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Albums found:\n")
	for _, album := range albums {
		fmt.Printf("ID=%d, Title=%s, Artist=%s, Price=%.2f\n", album.ID, album.Title, album.Artist, album.Price)
	}

	album, err := albumByID(db, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Album found :\nID=%d, Title=%s, Artist=%s, Price=%.2f\n", album.ID, album.Title, album.Artist, album.Price)

	albID, err := addAlbum(db, Album{Title: "The Modern Sound of Betty Carter", Artist: "Betty Carter", Price: 49.99})
	if err != nil {
		panic(err)
	}

	fmt.Printf("ID of added album: %v\n", albID)
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(db *gorm.DB, name string) ([]Album, error) {

	var albums []Album
	result := db.Find(&albums, "Artist = ?", name)
	return albums, result.Error
}

// albumByID queries for the album with the specified ID.
func albumByID(db *gorm.DB, id int64) (Album, error) {
	var alb Album
	result := db.First(&alb, id)
	return alb, result.Error
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(db *gorm.DB, alb Album) (int64, error) {
	result := db.Create(&alb)
	return alb.ID, result.Error
}
