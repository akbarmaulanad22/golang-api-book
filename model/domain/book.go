package domain

type Book struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Publisher  string `json:"publisher"`
	Year       int    `json:"year"`
	Genre      string `json:"genre"`
	Pages      int    `json:"pages"`
	CategoryId int    `json:"category_id"`
}