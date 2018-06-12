package controllers

import (
	"net/http"
)

/* BlogHandler serves the blog page */
func BlogHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.Get("blog").ExecuteTemplate(w, "base-tpl", nil)
	checkErr(err)
}
