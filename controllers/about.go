package controllers

import (
	"net/http"
)

/* ContactHandler serves the contact page */
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.Get("about").ExecuteTemplate(w, "base-tpl", nil)
	checkErr(err)
}
