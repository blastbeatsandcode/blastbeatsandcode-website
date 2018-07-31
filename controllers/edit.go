package controllers

import (
	"net/http"

	"github.com/blastbeatsandcode/blastbeatsandcode-website/utils"
)

/* ContactHandler serves the contact page */
func EditHandler(w http.ResponseWriter, r *http.Request) {
	isAuth := utils.HandleAccess(r)

	// Check if we have a username
	// If we do, tell the user they are logged in
	if isAuth {
		err := tpl.Get("edit").ExecuteTemplate(w, "base-tpl", nil)
		checkErr(err)
	} else {
		redirURL := "/"
		http.Redirect(w, r, redirURL, http.StatusSeeOther)
	}
}
