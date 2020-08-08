package main

import (
	"fmt"
	"github.com/AnuragManchanda/go-blog/app/controllers/articles"
	dbClient "github.com/AnuragManchanda/go-blog/db"
	"github.com/gorilla/mux"
	log "log"
	"net/http"
)

func main() {
	dbClient.Init()
	initializeRoutes()
}

func initializeRoutes() {
	fmt.Printf("Server Listening at port 80")
	router := mux.NewRouter()
	router.HandleFunc("/articles/{id}", articles.Show).Methods("GET")
	router.HandleFunc("/articles}", articles.Index).Methods("GET")
	router.HandleFunc("/articles}", articles.Create).Methods("POST")
	log.Fatal(http.ListenAndServe(":80", router))
}
