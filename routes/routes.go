package routes

import (
	"fmt"
	"net/http"

	"github.com/blastbeatsandcode/blastbeatsandcode-website/controllers"
	"github.com/gorilla/mux"
)

/* Routes returns the registered routes and handlers. */
func Routes() *mux.Router {
	// Create the new router
	r := mux.NewRouter().StrictSlash(false)

	// Create the subrouters and register controllers
	// We also will serve the public folders such as images, js, and css to each route

	// Register the Not Found catchall
	r.NotFoundHandler = http.HandlerFunc(notFound)

	// Register Home
	r.HandleFunc("/", controllers.HomeHandler)
	r.PathPrefix("/public/css/").Handler(http.StripPrefix("/public/css/", http.FileServer(http.Dir("./public/css/"))))
	r.PathPrefix("/public/js/").Handler(http.StripPrefix("/public/js/", http.FileServer(http.Dir("./public/js/"))))
	r.PathPrefix("/public/img/").Handler(http.StripPrefix("/public/img/", http.FileServer(http.Dir("./public/img/"))))

	// Register Blog
	blogRoute := r.PathPrefix("/blog").Subrouter()
	blogRoute.HandleFunc("", controllers.BlogHandler)
	blogRoute.PathPrefix("/public/css/").Handler(http.StripPrefix("/public/css/", http.FileServer(http.Dir("./public/css/"))))
	blogRoute.PathPrefix("/public/img/").Handler(http.StripPrefix("/public/img/", http.FileServer(http.Dir("./public/img/"))))
	blogRoute.PathPrefix("/public/js/").Handler(http.StripPrefix("/public/js/", http.FileServer(http.Dir("./public/js/"))))

	// Register Projects
	projectRoute := r.PathPrefix("/projects").Subrouter()
	projectRoute.HandleFunc("", controllers.ProjectsHandler)
	projectRoute.PathPrefix("/public/css/").Handler(http.StripPrefix("/public/css/", http.FileServer(http.Dir("./public/css/"))))
	projectRoute.PathPrefix("/public/img/").Handler(http.StripPrefix("/public/img/", http.FileServer(http.Dir("./public/img/"))))
	projectRoute.PathPrefix("/public/js/").Handler(http.StripPrefix("/public/js/", http.FileServer(http.Dir("./public/js/"))))

	// Register About
	aboutRoute := r.PathPrefix("/about").Subrouter()
	aboutRoute.HandleFunc("", controllers.AboutHandler)
	aboutRoute.PathPrefix("/public/css/").Handler(http.StripPrefix("/public/css/", http.FileServer(http.Dir("./public/css/"))))
	aboutRoute.PathPrefix("/public/img/").Handler(http.StripPrefix("/public/img/", http.FileServer(http.Dir("./public/img/"))))
	aboutRoute.PathPrefix("/public/js/").Handler(http.StripPrefix("/public/js/", http.FileServer(http.Dir("./public/js/"))))

	// Register Edit
	editRoute := r.PathPrefix("/edit").Subrouter()
	editRoute.HandleFunc("", controllers.EditHandler)
	editRoute.PathPrefix("/public/css/").Handler(http.StripPrefix("/public/css/", http.FileServer(http.Dir("./public/css/"))))
	editRoute.PathPrefix("/public/img/").Handler(http.StripPrefix("/public/img/", http.FileServer(http.Dir("./public/img/"))))
	editRoute.PathPrefix("/public/js/").Handler(http.StripPrefix("/public/js/", http.FileServer(http.Dir("./public/js/"))))

	// Register Login, this one has a POST and a GET because we use a form
	loginRoute := r.PathPrefix("/login").Subrouter()
	loginRoute.HandleFunc("", controllers.AuthGetHandler).Methods("GET")
	loginRoute.HandleFunc("", controllers.AuthPostHandler).Methods("POST")
	loginRoute.PathPrefix("/public/css/").Handler(http.StripPrefix("/public/css/", http.FileServer(http.Dir("./public/css/"))))
	loginRoute.PathPrefix("/public/img/").Handler(http.StripPrefix("/public/img/", http.FileServer(http.Dir("./public/img/"))))
	loginRoute.PathPrefix("/public/js/").Handler(http.StripPrefix("/public/js/", http.FileServer(http.Dir("./public/js/"))))

	// TODO: Add a logout route

	return r
}

/* notFound is the catchall handler for the gorilla router */
// TODO: create not found template and struct
func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Something broke!</h1>"+
		"<p>It looks like the page you were looking for has moved or never existed in the first place. :(</p>")
}
