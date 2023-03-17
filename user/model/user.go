package model

import (
	"time"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type (
	User struct {
		Id              int64     `xorm:"pk autoincr 'id'"`
		Username        string    `xorm:"notnull 'username'"`
		Password        string    `xorm:"notnull 'password'"`
		Enable          bool      `xorm:"default 1 'enable'"`             // 账号是否可用
		Deleted         bool      `xorm:"'deleted'"`                      // 账号删除(注销)标志位
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

	UserModel struct {
		db *xorm.Engine
		rdb *redis.Client
	}
)

func NewUserModel(DataSource string) *UserModel {
	mysqlDB, err := xorm.NewEngine("mysql", DataSource)
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
		return nil
	}
	mysqlDB.ShowSQL(true)
	model := &UserModel{
		db: mysqlDB,
		rdb: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			Password: "",
			DB: 0,
		}),
	}
	return model
}

func (m *UserModel) Insert(usr *User) {
	m.db.Cols("username", "password", "avatar", "background_image",
		"signature").Insert(usr)
}

func (m *UserModel) Exist(usr *User) (bool, error) {
	return m.db.Exist(usr)
}

func (m *UserModel) Update(usr *User) {
	m.db.ID(usr.Id).Update(*usr)
}

func (m *UserModel) Get(usr *User) (bool, error) {
	return m.db.Get(usr)
}