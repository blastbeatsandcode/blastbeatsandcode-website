package controllers

import (
	"fmt"
	"net/http"

	"github.com/blastbeatsandcode/blastbeatsandcode-website/utils"
)

// /* LoginHandler serves the login page */
// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	err := tpl.Get("login").ExecuteTemplate(w, "base-tpl", nil)
// 	checkErr(err)
// }

/* AuthGetHandler handles login GET requests */
func AuthGetHandler(w http.ResponseWriter, r *http.Request) {
	isAuth := utils.HandleAccess(r)

	// Check if we have a username
	// If we do, tell the user they are logged in
	if isAuth {
		redirURL := "/edit"
		http.Redirect(w, r, redirURL, http.StatusSeeOther)
	} else {
		err := tpl.Get("login").ExecuteTemplate(w, "base-tpl", nil)
		checkErr(err)
	}
}

/* Takes information from login POST requests and logs user in or shows error */
func AuthPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	err := utils.CheckLogin(username, password)

	if err != nil { // If error is not nil, then the login does not match
		username = ""
		fmt.Println("THERE WAS AN ERROR LOGGING IN")
	}

	// Save the session
	store := utils.GetStore()
	session, _ := store.Get(r, "session")

	session.Values["username"] = username
	session.Save(r, w)

	// Check if we have a matching project ID
	// If we do, load edit page, otherwise prompt for login
	isAuth := utils.HandleAccess(r)

	// If user is authorized to edit and the request matches, load edit page
	if isAuth {
		redirURL := "/edit"
		http.Redirect(w, r, redirURL, http.StatusSeeOther)
	} else { // Otherwise load the login-failed template
		session.Values["username"] = ""
		session.Save(r, w)
		err := tpl.Get("login").ExecuteTemplate(w, "base-tpl", nil)
		checkErr(err)
	}
}
