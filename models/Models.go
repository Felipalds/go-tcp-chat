package models

type User struct {
	Id   int
	Name string
}

type Room struct {
	Id       int
	Name     string
	Type     string
	Password string
	Admin    User
}

type Message struct {
	Id      int
	content string
	UserId  int
	RoomId  int
	Time    int
}
