package main

import (
	"net/http"
	"github.com/gorilla/mux"
	h "coconut.com/handlers"
	"github.com/gorilla/handlers"
	"os"
	"github.com/kabukky/httpscerts"
	"log"
)

func main() {
	// Generate cert
	// Check if the cert files are available.
	err := httpscerts.Check("cert.pem", "key.pem")
	// If they are not available, generate new ones.
	if err != nil {
		err = httpscerts.Generate("cert.pem", "key.pem", "127.0.0.1:8081")
		if err != nil {
			log.Fatal("Error: Couldn't create https certs.")
		}
	}

	// Here we are instantiating the gorilla/mux router
	r := mux.NewRouter()

	// On the default page we will simply serve our static index page.
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	// We will setup our server so we can serve static assest like images, css from the /static/{file} route
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.PathPrefix("/payloads/").Handler(http.StripPrefix("/payloads/", http.FileServer(http.Dir("./payloads/"))))

	r.Handle("/list", h.PayloadsHandler).Methods("GET")
	r.Handle("/upload", h.UploadHandler).Methods("POST")

	// Our application will run on port 8443. Here we declare the port and pass in our router.
	http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", handlers.LoggingHandler(os.Stdout, r))
	http.ListenAndServe(":8443", handlers.LoggingHandler(os.Stdout, r))
}
