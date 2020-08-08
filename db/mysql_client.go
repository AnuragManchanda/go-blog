package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {
	db, err := sql.Open("mysql", "dockeruser:dockeruser@tcp(mysqldb:33060)/blog?parseTime=true")
	//db, err := sql.Open("mysql", "dockeruser:dockeruser@tcp(localhost:33060)/blog?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database connected successfully")
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " +
		"articles (id VARCHAR(16) PRIMARY KEY, " +
		"title VARCHAR(16) NOT NULL, " +
		"author VARCHAR(16) NOT NULL, " +
		"content TEXT);")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Migration successfully Done")
	}
	DB = db
}

func Client() *sql.DB{
	return DB
}

//stmt, err := db.Prepare(“CREATE Table employee(id int NOT NULL AUTO_INCREMENT, first_name varchar(50), last_name varchar(30), PRIMARY KEY (id));”)
//if err != nil {
//fmt.Println(err.Error())
//}
//
//_, err = stmt.Exec()
//if err != nil {
//fmt.Println(err.Error())
//} else {
//fmt.Println(“Table created successfully..”)
//}
