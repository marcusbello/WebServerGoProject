package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Hello World!\n")
	if err != nil {
		return
	}
}

func main() {
	fmt.Println("Welcome to my GoServer")
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":3080", nil))
}
