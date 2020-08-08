package articles

import (
	"encoding/json"
	"fmt"
	"github.com/AnuragManchanda/go-blog/app/models/article"
	"net/http"
	"strconv"
	"strings"
)

type articleData struct {
	Title string
	Content  string
	Author string
}

func Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data articleData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	title := data.Title
	content := data.Content
	author := data.Author
	article.Create(title, author, content)
	fmt.Fprintf(w, "Successfully Created")
}

func Show(w http.ResponseWriter, r *http.Request) {
	id := strings.Split(r.RequestURI, "/")[2]
	id_n, _ := strconv.Atoi(id)
	response := article.Find(id_n)
	fmt.Fprintf(w, response)
}

func Index(w http.ResponseWriter, r *http.Request) {
	response := article.All()
	fmt.Fprintf(w, response)
}