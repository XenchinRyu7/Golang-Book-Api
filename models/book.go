package models

type Book struct {
	ID        int    `json:"id"`
	ISBN      string `json:"isbn"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Year      int    `json:"year"`
}
