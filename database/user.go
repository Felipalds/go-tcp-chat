package database

import (
	"fmt"
	"go-tcp-chat/models"
	"strings"
)

func GetUserById(id int64) ([]models.User, error) {
	db := GetDB()
	query := "SELECT * FROM users WHERE id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			fmt.Println(err)
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return users, nil
}

func GetUserByName(name string) (*models.User, error) {
	db := GetDB()
	query := "SELECT * FROM users WHERE name = ?"
	rows, err := db.Query(query, name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &user, nil
}

// Function to check if the error is due to a unique constraint violation
func isDuplicateKeyError(err error) bool {
	// This error message is specific to SQLite, adapt it for your database if needed
	return strings.Contains(err.Error(), "UNIQUE constraint failed")
}

func CreateNewUser(name string) (int64, error) {
	db := GetDB()
	query := "INSERT INTO users (name) values (?)"

	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(name)
	if err != nil {
		// Check if the error is due to a unique constraint violation
		if isDuplicateKeyError(err) {
			return 0, nil
		}
		fmt.Println(err)
		return 0, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	return lastId, err
}
