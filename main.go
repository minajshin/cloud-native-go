package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/minashin/cloud-native-go/api"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/echo", api.EchoHandlerFunc)
	http.HandleFunc("/api/books", api.BooksHandleFunc)
	http.HandleFunc("/api/books/", api.BookHandleFunc)
	http.ListenAndServe(port(), nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello Cloud Native Go!")
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
