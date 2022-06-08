package model

type Recipe struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Rating      int     `json:"rating"` 
}