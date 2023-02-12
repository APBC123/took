package models

type User struct {
	Id       int64
	Username string
	Password string
}

func (table User) TableName() string {
	return "user"
}
