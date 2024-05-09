package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" //não sei o que esse _ significa
)

func OpenDB() (*sql.DB, error) {
	// Opens of create DB
	// Check if the database file exists
	// To golang open files, it uses the GOPATH, which is the PATH of the .mod
	pwd, _ := os.Getwd()
	PATH := pwd + "/data/chat.db"
	_, err := os.Stat(PATH)
	if os.IsNotExist(err) {
		file, err := os.Create(PATH)
		if err != nil {
			fmt.Println("Error creating database")
		}
		file.Close()
	}

	// Open the database
	db, err := sql.Open("sqlite3", PATH)
	if err != nil {
		log.Fatal("Error opening database:", err)
		return nil, err
	}

	return db, nil
}

func Init(db *sql.DB) error {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (id integer not null primary key, name varchar(255))")
	if err != nil {
		log.Fatal("Error preparing statement:", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println("Error executing statement:", err)
	}
	return err
}
