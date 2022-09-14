package models

type Book struct {
	Id        int    `json:"id"`
	Judul     string `json:"judul"`
	Tahun     int    `json:"tahun"`
	Pengarang string `json:"pengarang"`
}
