package controllers

import (
	"fmt"
	"net/http"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Blast Beats and Code :: Blog</h1>"+
		"<p>The strange endeavors of a nerdy developer who also likes really bad music.</p>")
}
