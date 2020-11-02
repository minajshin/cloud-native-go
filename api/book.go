package api

import (
	"encoding/json"
	"net/http"
)

// Book defines Book type with Name, Author and ISBN
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

// Sample book data
var books = []Book{
	Book{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", ISBN: "0345391802"},
	Book{Title: "Cloud Native Go", Author: "M.-Leander Reimer", ISBN: "0000000000"},
}

// ToJSON marshals Book type to JSON
func (b Book) ToJSON() []byte {
	book, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return book
}

// FromJSON unmarshals JSON data to Book type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(books)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}
