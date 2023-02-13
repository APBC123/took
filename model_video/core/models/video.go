package models

type Video struct {
	Id            int64
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	Time          int
	Title         string
	CommentCount  int64
	FavoriteCount int64
}

func (table Video) TableName() string {
	return "video"
}
