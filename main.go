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
