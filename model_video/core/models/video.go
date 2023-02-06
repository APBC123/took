package models

type Video struct {
	Id       int
	AuthorId int
	PlayUrl  string
	CoverUrl string
	Time     int
	Title    string
	Removed  bool
	Deleted  bool
}

func (table Video) TableName() string {
	return "video"
}
