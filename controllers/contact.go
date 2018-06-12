package controllers

import (
	"net/http"
)

/* ContactHandler serves the contact page */
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.Get("contact").ExecuteTemplate(w, "base-tpl", nil)
	checkErr(err)
}
