package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.Handle("/form", http.HandlerFunc(formHandler))
	http.Handle("/hello", http.HandlerFunc(helloHandler))

	fmt.Println("Server Started at port: 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
