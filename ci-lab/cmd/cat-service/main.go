package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
  catHandler := func(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Aren't cats great?")
  }

  http.HandleFunc("/cat", catHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
