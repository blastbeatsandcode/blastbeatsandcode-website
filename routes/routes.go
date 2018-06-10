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

	// Register Contact
	contactRoute := r.PathPrefix("/contact").Subrouter()
	contactRoute.HandleFunc("", controllers.ContactHandler)
	contactRoute.PathPrefix("/public/css/").Handler(http.StripPrefix("/public/css/", http.FileServer(http.Dir("./public/css/"))))
	contactRoute.PathPrefix("/public/img/").Handler(http.StripPrefix("/public/img/", http.FileServer(http.Dir("./public/img/"))))
	contactRoute.PathPrefix("/public/js/").Handler(http.StripPrefix("/public/js/", http.FileServer(http.Dir("./public/js/"))))

	// Register Login
	loginRoute := r.PathPrefix("/login").Subrouter()
	loginRoute.HandleFunc("", controllers.LoginHandler)
	loginRoute.PathPrefix("/public/css/").Handler(http.StripPrefix("/public/css/", http.FileServer(http.Dir("./public/css/"))))
	loginRoute.PathPrefix("/public/img/").Handler(http.StripPrefix("/public/img/", http.FileServer(http.Dir("./public/img/"))))
	loginRoute.PathPrefix("/public/js/").Handler(http.StripPrefix("/public/js/", http.FileServer(http.Dir("./public/js/"))))

	return r
}

/* notFound is the catchall handler for the gorilla router */
// TODO: create not found template and struct
func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Something broke!</h1>"+
		"<p>It looks like the page you were looking for has moved or never existed in the first place. :(</p>")
}
