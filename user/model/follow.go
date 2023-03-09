package model

type Follow struct {
	UserId              int64     `xorm:"pk index 'user_id'"`
	FanId	int64 `xorm:"pk 'fan_id'"`
}
