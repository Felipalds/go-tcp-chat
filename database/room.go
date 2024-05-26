package database

import (
	"fmt"
	"go-tcp-chat/models"
)

func GetRoomByName(name string) (*models.Room, error) {
	db := GetDB()
	query := "SELECT rooms.id, rooms.name, rooms.type, rooms.password, users.id, users.name FROM rooms INNER JOIN users ON users.id = rooms.admin_id WHERE rooms.name = ?"
	rows, err := db.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var room models.Room
	var admin models.User
	for rows.Next() {
		if err := rows.Scan(&room.Id, &room.Name, &room.Type, &room.Password, &admin.Id, &admin.Name); err != nil {
			return nil, err
		}
		room.Admin = admin // Set the admin reference for the room
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &room, nil
}

func GetRooms() ([]models.Room, error) {
	db := GetDB()
	query := "SELECT rooms.id, rooms.name, rooms.type, rooms.password, users.id, users.name FROM rooms INNER JOIN users ON users.id = rooms.admin_id"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []models.Room
	var room models.Room
	var admin models.User
	for rows.Next() {
		if err := rows.Scan(&room.Id, &room.Name, &room.Type, &room.Password, &admin.Id, &admin.Name); err != nil {
			return nil, err
		}
		room.Admin = admin // Set the admin reference for the room
		rooms = append(rooms, room)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return rooms, nil
}
func CreateNewRoom(room models.Room) (int64, error) {
	//TODO: tentar reduzir este arquivo
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

func AddUserToRoom(userID int64, roomID int64) error {
	db := GetDB()

	query := "INSERT INTO user_room (user_id, room_id, is_banned) VALUES (?, ?, ?)"

	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, roomID, false)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
