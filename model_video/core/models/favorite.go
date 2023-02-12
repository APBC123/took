package models

type Favorite struct {
	Id      int64
	VideoId int64
	UserId  int64
}

func (table Favorite) TableName() string {
	return "favorite"
}
