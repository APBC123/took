package handler

import (
	"bufio"
	"encoding/json"
	"github.com/zeromicro/go-zero/rest/httpx"
	"model_video/core/helper"
	"model_video/core/internal/logic"
	"model_video/core/internal/svc"
	"model_video/core/internal/types"
	"net/http"
	"os"
	"path"
)

func PublishVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		//获取上传文件
		file, fileHeader, err := r.FormFile("data")
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			var ret, _ = json.Marshal(&types.PublishResponse{Response: types.Response(struct {
				StatusCode int32
				StatusMsg  string
			}{StatusCode: 1, StatusMsg: "get file failed"})})
			w.Header().Set("Content-Type", "application-json")
			w.Write(ret)
			return
		}
		fv, err := os.OpenFile("./video/"+helper.UUID()+path.Ext(fileHeader.Filename), os.O_CREATE|os.O_RDONLY, 0777)
		if err != nil {
			var ret, _ = json.Marshal(&types.PublishResponse{Response: types.Response(struct {
				StatusCode int32
				StatusMsg  string
			}{StatusCode: 2, StatusMsg: "get file failed"})})
			w.Header().Set("Content-Type", "application-json")
			w.Write(ret)
			return
		}
		wf := bufio.NewWriter(fv)
		wf.Write(b)
		wf.Flush()
		fv.Close()
		//生成截图
		err, filename := helper.GetVideoShot(fv.Name(), "./cover/"+helper.UUID(), 1)
		if err != nil {
			var ret, _ = json.Marshal(&types.PublishResponse{Response: types.Response(struct {
				StatusCode int32
				StatusMsg  string
			}{StatusCode: 3, StatusMsg: "get file failed"})})
			w.Header().Set("Content-Type", "application-json")
			w.Write(ret)
			return
		}

		//储存截图
		fc, err := os.OpenFile("./cover/"+filename, os.O_RDONLY, 0666)
		if err != nil {
			var ret, _ = json.Marshal(&types.PublishResponse{Response: types.Response(struct {
				StatusCode int32
				StatusMsg  string
			}{StatusCode: 4, StatusMsg: "get file failed"})})
			w.Header().Set("Content-Type", "application-json")
			w.Write(ret)
			return
		}
		coverPath, err := helper.CosUploadLocal(fc, filename)
		if err != nil {
			var ret, _ = json.Marshal(&types.PublishResponse{Response: types.Response(struct {
				StatusCode int32
				StatusMsg  string
			}{StatusCode: 5, StatusMsg: "upload file failed"})})
			w.Header().Set("Content-Type", "application-json")
			w.Write(ret)
		}

		//向腾讯云cos中储存文件
		videoPath, err := helper.CosUpload(r)
		if err != nil {
			var ret, _ = json.Marshal(&types.PublishResponse{Response: types.Response(struct {
				StatusCode int32
				StatusMsg  string
			}{StatusCode: 6, StatusMsg: "upload file failed"})})
			w.Header().Set("Content-Type", "application-json")
			w.Write(ret)
		}

		req.PlayUrl = videoPath
		req.CoverUrl = coverPath
		//helper.GetVideoShot

		l := logic.NewPublishVideoLogic(r.Context(), svcCtx)
		resp, err := l.PublishVideo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
