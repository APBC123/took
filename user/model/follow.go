package model

import (
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type (
	Follow struct {
		UserId int64 `xorm:"pk index 'user_id'"`
		FanId  int64 `xorm:"pk 'fan_id'"`
	}

	FollowModel struct {
		db *xorm.Engine
		rdb *redis.Client
	}
)

func NewFollowModel(DataSource string) *FollowModel {
	mysqlDB, err := xorm.NewEngine("mysql", DataSource)
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
		return nil
	}
	mysqlDB.ShowSQL(true)
	model := &FollowModel{
		db: mysqlDB,
		rdb: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			Password: "",
			DB: 0,
		}),
	}
	return model
}

func (m *FollowModel) Insert(follow *Follow) {
	m.db.Insert(follow)
}

func (m *FollowModel) Exist(follow *Follow) (bool, error){
	return m.db.Exist(follow)
}

func (m *FollowModel) Delete(follow *Follow) {
	m.db.Delete(follow)
}

func (m *FollowModel) FindFollowerById(id int64) (followerList []*User) {
	m.db.Table("user").Join("LEFT", "follow", "user.id = follow.fan_id").Select(
		"user.*").Where("follow.user_id = ?", id).Find(&followerList)
	return followerList
}

func (m *FollowModel) FindFollowById(id int64) (followList []*User) {
	m.db.Table("user").Join("LEFT", "follow", "user.id = follow.user_id").Select(
		"user.*").Where("follow.fan_id = ?", id).Find(&followList)
	return followList
}

func (m *FollowModel) FindFriendById(id int64) (friendList []*User) {
	m.db.Table("user").Alias("u").Join(
		"INNER", []string{"follow", "f1"}, "u.id = f1.user_id").Join(
		"INNER", []string{"follow", "f2"}, "f1.user_id = f2.fan_id").Select("u.*").Where(
		"f1.fan_id = ? AND f1.fan_id = f2.user_id", id).Find(&friendList)
	return friendList
}