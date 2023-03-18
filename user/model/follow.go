package model

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

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

func (m *FollowModel) Exist(ctx context.Context, follow *Follow) (bool, error){
	key := "follow:"+strconv.FormatInt(follow.FanId, 10)+":"+strconv.FormatInt(follow.UserId, 10)+":exist"
	strHas, err := m.rdb.Get(ctx, key).Result()
	if err == nil {
		log.Printf("get key <- %s = %s\n", key, strHas)
		return strconv.ParseBool(strHas)
	}
	has, _ := m.db.Exist(follow)
	m.rdb.Set(ctx, key, has, 0)
	log.Printf("set key -> %s = %v\n", key, has)
	return has, nil	
}

func (m *FollowModel) Delete(follow *Follow) {
	m.db.Delete(follow)
}

func (m *FollowModel) GetFollowerById(ctx context.Context, id int64) (followerList []*User) {
	key := "followerList:"+strconv.FormatInt(id, 10)
	followerListStr, err := m.rdb.Get(ctx, key).Result()
	if err == nil {
		log.Printf("get key <- %s = %s\n", key, followerListStr)
		err = json.Unmarshal([]byte(followerListStr), &followerList)
		if err != nil {
			return nil
		}
		return followerList
	}

	m.db.Table("user").Join("LEFT", "follow", "user.id = follow.fan_id").Select(
		"user.*").Where("follow.user_id = ?", id).Find(&followerList)
	followerListData, _ := json.Marshal(followerList)
	m.rdb.Set(ctx, key, followerListData, 0)
	log.Printf("set key -> %s = %s\n", key, followerListData)
	return followerList
}

func (m *FollowModel) GetFollowById(ctx context.Context, id int64) (followList []*User) {
	key := "followList:"+strconv.FormatInt(id, 10)
	followListStr, err := m.rdb.Get(ctx, key).Result()
	if err == nil {
		log.Printf("get key <- %s = %s\n", key, followListStr)
		err = json.Unmarshal([]byte(followListStr), &followList)
		if err != nil {
			return nil
		}
		return followList
	}

	m.db.Table("user").Join("LEFT", "follow", "user.id = follow.user_id").Select(
		"user.*").Where("follow.fan_id = ?", id).Find(&followList)
	followListData, _ := json.Marshal(followList)
	m.rdb.Set(ctx, key, followListData, 0)
	log.Printf("set key -> %s = %s\n", key, followListData)
	return followList
}

func (m *FollowModel) GetFriendById(ctx context.Context, id int64) (friendList []*User) {
	key := "friendList:"+strconv.FormatInt(id, 10)
	friendListStr, err := m.rdb.Get(ctx, key).Result()
	if err == nil {
		log.Printf("get key <- %s = %s\n", key, friendListStr)
		err = json.Unmarshal([]byte(friendListStr), &friendList)
		if err != nil {
			return nil
		}
		return friendList
	}

	m.db.Table("user").Alias("u").Join(
		"INNER", []string{"follow", "f1"}, "u.id = f1.user_id").Join(
		"INNER", []string{"follow", "f2"}, "f1.user_id = f2.fan_id").Select("u.*").Where(
		"f1.fan_id = ? AND f1.fan_id = f2.user_id", id).Find(&friendList)
	friendListData, _ := json.Marshal(friendList)
	m.rdb.Set(ctx, key, friendListData, 0)
	log.Printf("set key -> %s = %s\n", key, friendListData)
	return friendList
}