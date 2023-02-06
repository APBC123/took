package models

import "time"

type User struct {
	Id          int
	Username    string
	Password    string
	Enable      bool
	Deleted     bool
	LoginTime   time.Time `xorm:"updated"`
	CreatedTime time.Time `xorm:"created"`
}
