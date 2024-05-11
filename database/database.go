package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	_ "embed"
	_ "github.com/mattn/go-sqlite3" //não sei o que esse _ significa
)

// SINGLETON!!! :D
var lock = &sync.Mutex{}
var databaseInstance *sql.DB

//go:embed database.sql
var databaseQuery string

//go:embed trigger.sql
var trigger string

func GetDB() *sql.DB {
	if databaseInstance == nil {
		lock.Lock()
		db, _ := OpenDB()
		databaseInstance = db
		lock.Unlock()
	}
	return databaseInstance
}

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

	queries := strings.Split(databaseQuery, ";")
	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}
		_, err := db.Exec(query)
		if err != nil {
			return fmt.Errorf("Error executing query: %v", err)
		}
	}

	trigg := strings.TrimSpace(trigger)
	_, err := db.Exec(trigg)
	fmt.Println(trigg)
	if err != nil {
		return fmt.Errorf("Error executing trigger: %v", err)
	}

	fmt.Println("Database initialized")
	return nil
}
