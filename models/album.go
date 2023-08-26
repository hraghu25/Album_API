package models

type Album struct {
	ID     int    `json:"id"`
	Title  string `json:"string"`
	Artist string `json:"artist"`
	Price  string `json:"price"`
}

type AlbumSlice []Album

var Albums = AlbumSlice{}
