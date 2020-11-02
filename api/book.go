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

// in-memory data structure
var books = map[string]Book{
	"0345391802": Book{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", ISBN: "0345391802"},
	"0000000000": Book{Title: "Cloud Native Go", Author: "M.-Leander Reimer", ISBN: "0000000000"},
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

func writeJson(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

// AllBooks returns an array of books of database
func AllBooks() []Book {
	values := make([]Book, len(books))
	i := 0
	for _, book := range books {
		values[i] = book
		i++
	}
	return values
}

// BooksHandleFunc sends a list of all books
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJson(w, books)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method"))
	}
}

// Books
