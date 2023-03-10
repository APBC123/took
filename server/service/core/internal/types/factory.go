package types

import "took/user/rpc/types/user"

func NewUser(usr *user.User) (User) {
	return User{
		Id: usr.Id,
		Username: usr.Username,
		FollowCount: usr.FollowCount,
		FollowerCount: usr.FollowerCount,
		IsFollow: usr.IsFollow,
		Avatar: usr.Avatar,
		BackgroundImage: usr.BackgroundImage,
		Signature: usr.Signature,
		TotalFavorited: usr.TotalFavorited,
		WorkCount: usr.WorkCount,
		FavoriteCount: usr.FavoriteCount,
	}
}

func NewUserList(usrList []*user.User) ([]User) {
	respList := make([]User, len(usrList))
	for i := range usrList {
		respList[i] = NewUser(usrList[i])
	}
	return respList
}

func NewFriendUser(usr *user.FriendUser) (FriendUser) {
	return FriendUser{
		Id: usr.Id,
		Username: usr.Username,
		FollowCount: usr.FollowCount,
		FollowerCount: usr.FollowerCount,
		IsFollow: usr.IsFollow,
		Avatar: usr.Avatar,
		BackgroundImage: usr.BackgroundImage,
		Signature: usr.Signature,
		TotalFavorited: usr.TotalFavorited,
		WorkCount: usr.WorkCount,
		FavoriteCount: usr.FavoriteCount,
		Message: usr.Message,
		MsgType: usr.MsgType,
	}
}

func NewFriendUserList(usrList []*user.FriendUser) ([]FriendUser) {
	respList := make([]FriendUser, len(usrList))
	for i := range usrList {
		respList[i] = NewFriendUser(usrList[i])
	}
	return respList
}