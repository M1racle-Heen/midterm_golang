package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/store", returnAllArticles)
	myRouter.HandleFunc("/store/{id}", Get)
	myRouter.HandleFunc("/store/{id}/{item}", Put)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
