package controllers

import (
	"net/http"
)

/* HomeHandler serves the index page */
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.Get("index").ExecuteTemplate(w, "base-tpl", nil)
	checkErr(err)
}
