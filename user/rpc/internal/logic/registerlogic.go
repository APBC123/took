package logic

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"took/user/model"
	"took/user/rpc/internal/svc"
	"took/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	has, _ := l.svcCtx.Engine.Exist(&model.User{
		Username: in.Username,
	})
	if has {
		return &user.RegisterResp{
			StatusCode: 1,
			StatusMsg:  "用户名已存在",
		}, nil
	}

	// 随机生成用户头像、首页背景、个性签名
	var text struct {
		Code    string `json:"code"`
		Type    string `json:"type"`
		Content string `json:"content"`
	}
	for {
		resp, err := http.Get("https://api.uixsj.cn/hitokoto/get?type=hitokoto&code=json")
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)

		json.Unmarshal(body, &text)
		if len(text.Content) < 90 { 
			break;
		}
		time.Sleep(30*time.Millisecond)
	}
	usr := model.User{
		Username:        in.Username,
		Password:        fmt.Sprintf("%x", md5.Sum([]byte(in.Password))), // 哈希加密
		Avatar:          "https://www.loliapi.com/acg/pp?id=" + fmt.Sprint(rand.Intn(210)),
		BackgroundImage: "https://www.loliapi.com/acg/pc/?id=" + fmt.Sprint(rand.Intn(700)),
		Signature:       text.Content,
	}

	l.svcCtx.Engine.Cols("username", "password", "avatar", "background_image",
		"signature").Insert(&usr)

	return &user.RegisterResp{
		StatusCode: 0,
		StatusMsg:  "注册成功",
		UserId:     usr.Id,
	}, nil
}
