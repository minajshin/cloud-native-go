package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "M. -L. Reimer", ISBN: "0123456789"}
	json := book.ToJSON()

	assert.Equal(t, `{"Title":"Cloud Native Go","Author":"M. -L. Reimer","ISBN":"0123456789"}`,
		string(json), "Book json marshalling wrong")
}
