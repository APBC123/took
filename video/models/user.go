package models

type User struct {
	Id              int64
	Username        string
	Password        string
	FollowCount     int64
	FollowerCount   int64
	Avatar          string
	BackgroundImage string
	Signature       string
}

func (table User) TableName() string {
	return "user"
}
