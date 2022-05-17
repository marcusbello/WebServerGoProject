package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func hellHandler(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Hell World!\n")
	if err != nil {
		return
	}
}

func main() {
	fmt.Println("Welcome to my GoServer")

	http.HandleFunc("/", hellHandler)
	log.Fatal(http.ListenAndServe(":3080", nil))
}
