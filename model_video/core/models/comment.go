package models

type Comment struct {
	Id      int64
	UserId  int64
	VideoId int64
	Content string
}

func (table Comment) TableName() string {
	return "comment"
}
