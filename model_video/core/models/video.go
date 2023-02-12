package models

type Video struct {
	Id       int64
	Author   int64
	PlayUrl  string
	CoverUrl string
	Time     int
	Title    string
}

func (table Video) TableName() string {
	return "video"
}
