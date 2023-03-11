package models

import "time"

type Message struct {
	Id         string
	SendId     int64
	RecvId     int64
	Content    string
	CreateTime time.Time
}

func (table Message) TableName() string {
	return "message"

}
