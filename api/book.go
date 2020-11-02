package api

import (
	"encoding/json"
)

type Book struct {
	Title  string
	Author string
	ISBN   string
}

func (b Book) ToJSON() []byte {
	book, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return book
}

func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}
