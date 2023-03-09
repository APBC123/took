package helper

import (
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
