package articles

import (
	"fmt"
	"github.com/AnuragManchanda/go-blog/app/models/article"
	"net/http"
)

func Show(w http.ResponseWriter, r *http.Request) {
	response := article.Find(1)
	fmt.Fprintf(w, response)
}

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Print("hello2")
	article.Create()
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Print("hello3")
	article.Create()
}