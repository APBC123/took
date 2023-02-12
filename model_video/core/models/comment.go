package models

import "time"

type Comment struct {
	UserId     int64
	VideoId    int64
	Content    string
	CreateTime time.Time
}

func (table Comment) TableName() string {
	return "comment"
}
