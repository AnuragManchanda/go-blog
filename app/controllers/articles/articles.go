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
	//id := strings.Split(r.RequestURI, "/")[2]
	//id_n, _ := strconv.Atoi(id)
	//response := article.Find(id_n)
	//fmt.Fprintf(w, response)
}

func Index(w http.ResponseWriter, r *http.Request) {
	//	response := article.All()
	//	fmt.Fprintf(w, response)
	//}
}