package controllers

import (
	"net/http"
)

/* ProjectsHandler serves the projects page */
func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.Get("projects").ExecuteTemplate(w, "base-tpl", nil)
	checkErr(err)
}
