package articles

import (
	"fmt"
	"github.com/AnuragManchanda/go-blog/app/models/article"
	"net/http"
	"strconv"
	"strings"
)

func Show(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.RequestURI, "/")[2]
	id_n, _ := strconv.Atoi(id)
	response := article.Find(id_n)
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