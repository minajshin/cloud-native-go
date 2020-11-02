package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "M. -L. Reimer", ISBN: "0123456789"}
	json := book.ToJSON()

	assert.Equal(t, `{"title":"Cloud Native Go","author":"M. -L. Reimer","isbn":"0123456789"}`,
		string(json), "Book json marshalling wrong")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"Cloud Native Go","author":"M.-L. Reimer","isbn":"0123456789"}`)
	book := FromJSON(json)

	assert.Equal(t, Book{Title: "Cloud Native Go", Author: "M.-L. Reimer", ISBN: "0123456789"},
		book, "Book JSON unmarshalling wrong.")
}
