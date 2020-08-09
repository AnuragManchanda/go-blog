package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var DB *sql.DB

func Init() {
	for true {
		db, err := sql.Open("mysql", "dockeruser:dockeruser@tcp(mysqldb:3306)/blog?parseTime=true")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Database connected successfully")
		}
		result, err := db.Exec("CREATE TABLE IF NOT EXISTS " +
			"articles" +
			"(id int PRIMARY KEY AUTO_INCREMENT, " +
			"title VARCHAR(16) NOT NULL, " +
			"author VARCHAR(16) NOT NULL, " +
			"content TEXT);")
		if err != nil {
			fmt.Println(err.Error())
			time.Sleep(2 * time.Second)
		} else {
			fmt.Print(result)
			fmt.Println("Migration successfully Done")
			DB = db
			break
		}
	}
}

func Client() *sql.DB{
	return DB
}

func SetTestDb(db *sql.DB){
	DB = db
}
