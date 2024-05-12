package models

type User struct {
	Id   int64
	Name string
}

type Room struct {
	Id       int64
	Name     string
	Type     string
	Password string
	Admin    User
}

type Message struct {
	Id      int64
	content string
	UserId  int64
	RoomId  int64
	Time    int64
}
