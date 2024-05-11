package tests

import (
	"go-tcp-chat/database"
	"os"
	"testing"
)

func TestDatabaseCreation(t *testing.T) {

	os.Remove("./data/chat.db")

	db, err := database.OpenDB()
	if err != nil {
		t.Errorf("Error creating database %v", err)
	}

	err = database.Init(db)
	if err != nil {
		t.Errorf("Error initing database %v", err)
	}
}
