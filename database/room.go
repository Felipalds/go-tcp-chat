package database

import (
	"fmt"
	"go-tcp-chat/models"
)

func GetRoomByName(name string) (*models.Room, error) {
	db := GetDB()
	query := "SELECT * FROM rooms WHERE name = ? INNER JOIN admin ON admin.id = rooms.admin_id"
	rows, err := db.Query(query, name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var room models.Room
	for rows.Next() {
		if err := rows.Scan(&room.Id, &room.Name, &room.Password, &room.Admin.Id, &room.Admin.Name); err != nil {
			fmt.Println(err)
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &room, nil
}

func CreateNewRoom(room models.Room) (int64, error) {
	db := GetDB()

	query := "INSERT INTO rooms (name, type, password, admin_id) values (?, ?, ?, ?)"

	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(room.Name, room.Type, room.Password, room.Admin.Id)
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
