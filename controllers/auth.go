package controllers

import (
	"net/http"
)

/* LoginHandler serves the login page */
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.Get("login").ExecuteTemplate(w, "base-tpl", nil)
	checkErr(err)
}
