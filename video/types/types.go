package types

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	Title         string `json:"title,omitempty"`
	IsFavorite    bool   `json:"is_favorite"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Username      string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

type FeedRequest struct {
	LastTime int64  `json:"latest_time,optional"`
	Token    string `json:"token,optional"`
}

type PublishRequest struct {
	Token    string `json:"token,optional"`
	Title    string `json:"title,optional"`
	PlayUrl  string `json:"play_url,optional"`
	CoverUrl string `json:"cover_url,optional"`
}

type PublishResponse struct {
	Response
}

type PublishListRequest struct {
	Token  string `json:"token"`
	UserId int64  `json:"user_id"`
}

type PublishListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

type CommentListRequest struct {
	Token   string `json:"token,omitempty"`
	VideoId int64  `json:"video_id,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

type CommentListResponse struct {
	StatusCode  int32     `json:"status_code"`
	StatusMsg   string    `json:"status_msg,optional"`
	CommentList []Comment `json:"comment_list,omitempty"`
}
