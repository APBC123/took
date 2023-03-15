package helper

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
	"took/chat/rpc/types/chat"
	"took/server/service/core/internal/types"
	"took/video/video/rpc/types/video"
)

func NewUser(user *video.User) types.User {
	return types.User{
		Id:              user.Id,
		Username:        user.Username,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
	}
}

func NewCommentList(comment []*video.Comment) []types.Comment {
	list := make([]types.Comment, len(comment))
	for i := range list {
		list[i].Id = comment[i].Id
		list[i].Content = comment[i].Content
		list[i].CreateDate = comment[i].CreateDate
		list[i].User = NewUser(comment[i].User)
	}
	return list
}

func NewVideoList(video []*video.Video) []types.Video {
	list := make([]types.Video, len(video))
	for i := range list {
		list[i].Id = video[i].Id
		list[i].FavoriteCount = video[i].FavoriteCount
		list[i].IsFavorite = video[i].IsFavorite
		list[i].CommentCount = video[i].CommentCount
		list[i].Title = video[i].Title
		list[i].CoverUrl = video[i].CoverUrl
		list[i].PlayUrl = video[i].PlayUrl
		list[i].Author = NewUser(video[i].Author)
	}
	return list
}

func NewChatMessageList(chatMessage []*chat.Message) []types.Message {
	list := make([]types.Message, len(chatMessage))
	for i := range list {
		list[i].ToUserId = chatMessage[i].ToUserId
		list[i].Id = chatMessage[i].Id
		list[i].Content = chatMessage[i].Content
		list[i].FromUserId = chatMessage[i].FromUserId
		list[i].CreateTime = chatMessage[i].CreateTime
	}
	return list
}

func Random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(300)
}

var JwtKey = "took"

func AnalyzeTokenN(token string) (*UserClaim, error) {
	uc := new(UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
}
