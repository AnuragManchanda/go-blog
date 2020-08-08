package article

import (
	"encoding/json"
	"fmt"
	dbClient "github.com/AnuragManchanda/go-blog/db"
)

//type ArticleRow struct {
//	Id int
//	Title  string
//	Author  string
//	Content  string
//}
//
//type Article struct {
//	Id int 			`json:"id"`
//	Title  string   `json:"title`
//	Author  string  `json:"author`
//	Content  string `json:"content`
//}

type Article struct {
	Id int
	Title  string
	Author  string
	Content  string
}

type Articles struct {
	Articles []Article
}

func Find(id int) string {
	// Execute the query
	results, err := dbClient.Client().Query("SELECT id, title, author, content FROM ARTICLES")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var articles []Article

	for results.Next() {
		var articleRow Article
		// for each row, scan the result into our tag composite object
		err = results.Scan(&articleRow.Id, &articleRow.Title, &articleRow.Author, &articleRow.Content)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		article := Article{articleRow.Id,articleRow.Title,articleRow.Author,articleRow.Content}
		articles = append(articles, article)
		//articles := []Article{Article{1,"24","as","sd"}, Article{2,"234","as","sd"}}
	}
	articlesSlice, err := json.Marshal(Articles{articles})
	if err != nil {
		fmt.Println("Error...")
	}
	fmt.Println(string(articlesSlice))
	return string(articlesSlice)
}

func All() []byte {
	// Execute the query
	results, err := dbClient.Client().Query("SELECT id, title, author, content FROM ARTICLES")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var articles []Article

	for results.Next() {
		var articleRow Article
		// for each row, scan the result into our tag composite object
		err = results.Scan(&articleRow.Id, &articleRow.Title, &articleRow.Author, &articleRow.Content)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		article := Article{articleRow.Id,articleRow.Title,articleRow.Author,articleRow.Content}
		articles = append(articles, article)
		//articles := []Article{Article{1,"24","as","sd"}, Article{2,"234","as","sd"}}
	}
	articlesSlice, err := json.Marshal(Articles{articles})
	if err != nil {
		fmt.Println("Error...")
	}
	fmt.Println(string(articlesSlice))
	return articlesSlice
}


func Create() bool {
	return true
}