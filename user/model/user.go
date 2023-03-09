package model

import "time"

type User struct {
	Id              int64     `xorm:"pk autoincr 'id'"`
	Username        string    `xorm:"notnull 'username'"`
	Password        string    `xorm:"notnull 'password'"`
	Enable          int64     `xorm:"default 1 'enable'"`                       // 账号是否可用
	Deleted         int64     `xorm:"'deleted'"`                      // 账号删除(注销)标志位
	LoginTime       time.Time `xorm:"datetime 'login_time'"`          // 最近一次登录的时间
	CreateTime      time.Time `xorm:"datetime created 'create_time'"` // 注册时间
	FollowCount     int64     `xorm:"'follow_count'"`                 // 关注数
	FollowerCount   int64     `xorm:"'follower_count'"`               // 粉丝数
	Avatar          string    `xorm:"default Null 'avatar'"`          // 用户头像
	BackgroundImage string    `xorm:"'background_image'"`             // 用户个人页顶部大图url
	Signature       string    `xorm:"'signature'"`                    // 用户简介
	TotalFavorited  int64     `xorm:"'total_favorited'"`              // 用户获赞数量
	WorkCount       int64     `xorm:"'work_count'"`                   // 用户作品数
	FavoriteCount   int64     `xorm:"'favorite_count'"`               // 用户喜欢的作品数
}
