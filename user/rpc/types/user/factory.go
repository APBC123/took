package user

import "took/user/model"

func NewUser(usr *model.User) (*User) {
	return &User{
		Id: usr.Id,
		Username: usr.Username,
		FollowCount: usr.FollowCount,
		FollowerCount: usr.FollowerCount,
		Avatar: usr.Avatar,
		BackgroundImage: usr.BackgroundImage,
		Signature: usr.Signature,
		TotalFavorited: usr.TotalFavorited,
		WorkCount: usr.WorkCount,
		FavoriteCount: usr.FavoriteCount,
	}
}

func NewUserList(usrList []*model.User) ([]*User) {
	respList := make([]*User, len(usrList))
	for i := range usrList {
		respList[i] = NewUser(usrList[i])
	}
	return respList
}

func NewFriend(usr *model.User) (*FriendUser) {
	return &FriendUser{
		Id: usr.Id,
		Username: usr.Username,
		FollowCount: usr.FollowCount,
		FollowerCount: usr.FollowerCount,
		Avatar: usr.Avatar,
		BackgroundImage: usr.BackgroundImage,
		Signature: usr.Signature,
		TotalFavorited: usr.TotalFavorited,
		WorkCount: usr.WorkCount,
		FavoriteCount: usr.FavoriteCount,
	}
}

func NewFriendList(usrList []*model.User) ([]*FriendUser) {
	respList := make([]*FriendUser, len(usrList))
	for i := range usrList {
		respList[i] = NewFriend(usrList[i])
	}
	return respList
}