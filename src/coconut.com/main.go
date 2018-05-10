package main

import (
	"net/http"
	"github.com/gorilla/mux"
	h "coconut.com/handlers"
	"github.com/gorilla/handlers"
	"os"
)

func main() {
	// Here we are instantiating the gorilla/mux router
	r := mux.NewRouter()

	// On the default page we will simply serve our static index page.
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	// We will setup our server so we can serve static assest like images, css from the /static/{file} route
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.PathPrefix("/payloads/").Handler(http.StripPrefix("/payloads/", http.FileServer(http.Dir("./payloads/"))))

	r.Handle("/list", h.PayloadsHandler).Methods("GET")
	r.Handle("/upload", h.UploadHandler).Methods("POST")
	r.Handle("/event_handler", h.EventHandler).Methods("POST")
	r.Handle("/build", h.BuildHandler).Methods("POST")

	// Our application will run on port 8443. Here we declare the port and pass in our router.
	http.ListenAndServe(":4000", handlers.LoggingHandler(os.Stdout, r))
}
