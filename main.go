package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Blast Beats and Code</h1>"+
		"<p>The strange endeavors of a nerdy computer scientist who also likes really bad music.</p>")
}

func main() {

	// Register and handle home
	http.HandleFunc("/", handler)
	fmt.Println("Starting Server...")

	// Serve Project
	log.Fatal(http.ListenAndServe(":8080", nil))
}
