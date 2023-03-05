package models

type Video struct {
	Id            int64
	AuthorId      int64
	PlayUrl       string
	CoverUrl      string
	Title         string
	CommentCount  int64
	FavoriteCount int64
}

func (table Video) TableName() string {
	return "video"
}
