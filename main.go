package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	serverRouter := mux.NewRouter()
	serverRouter.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	serverRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/home.html")
	})
	serverRouter.HandleFunc("/tail/auth", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/spotify-return.html")
	})
	serverRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "views/construction.html")
	})
	err := http.ListenAndServe(":8080", serverRouter)
	if err != nil {
		log.Fatal("There was an error starting the server: ", err)
	}
}
