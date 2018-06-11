package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/blastbeatsandcode/blastbeatsandcode-website/routes"
)

func main() {

	// Register and handle home
	r := routes.Routes()

	// Print out something to the screen so we know that the server is starting
	fmt.Println("Starting Server...")

	// Connect to database and serve project
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
