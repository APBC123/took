package models

type Follow struct {
	Id     int
	UserId int64
	FanId  int64
}

func (table Follow) TableName() string {
	return "follow"
}
