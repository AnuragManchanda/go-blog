package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>You can Do it!</h1>")
}

func main() {
	fmt.Printf("Server Listening at port 80")
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":80", nil)
}