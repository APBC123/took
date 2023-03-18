package model

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"

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

func (m *UserModel) Exist(ctx context.Context, usr *User) (bool, error) {
	key := "user:name:"+usr.Username+":exist"
	strHas, err := m.rdb.Get(ctx, key).Result()
	if err == nil {
		log.Printf("get key <- %s = %s\n", key, strHas)
		return strconv.ParseBool(strHas)
	}
	has, _ := m.db.Exist(usr)
	m.rdb.Set(ctx, key, has, 0)
	log.Printf("set key -> %s = %v\n", key, has)
	return has, nil
}

func (m *UserModel) Update(ctx context.Context, usr *User) {
	m.db.ID(usr.Id).Update(*usr)
}

func (m *UserModel) GetById(ctx context.Context, usr *User) error {
	key := "user:id:"+strconv.FormatInt(usr.Id, 10)
	userStr, err := m.rdb.Get(ctx, key).Result()
	if err == nil {
		log.Printf("get key <- %s = %s\n", key, userStr)
		err = json.Unmarshal([]byte(userStr), &usr)
		if err != nil {
			return err
		}
		return nil
	}
	
	_, err = m.db.ID(usr.Id).Get(usr)
	if err != nil {
		return err
	}
	userData, _ := json.Marshal(usr)
	m.rdb.Set(ctx, key, userData, 0)
	log.Printf("set key -> %s = %s\n", key, userData)
	return nil
}

func (m *UserModel) GetByName(ctx context.Context, usr *User) (bool, error) {
	key := "user:name:"+usr.Username
	has, err := m.Exist(ctx, usr)
	if err != nil || !has {
		return has, err
	}
	userStr, err := m.rdb.Get(ctx, key).Result()
	if err == nil {
		log.Printf("get key <- %s = %s\n", key, userStr)
		err = json.Unmarshal([]byte(userStr), &usr)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	_, err = m.db.Where("username=?", usr.Username).Get(usr)
	if err != nil {
		return false, err
	}
	userData, _ := json.Marshal(usr)
	m.rdb.Set(ctx, key, userData, 0)
	log.Printf("set key -> %s = %s\n", key, userData)
	return true, nil
}