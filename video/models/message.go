package models

type Message struct {
	Id         string `xorm:"autoincr"`
	FromUserId int64
	ToUserId   int64
	Content    string
	CreateTime int64
}

func (table Message) TableName() string {
	return "message"

}
