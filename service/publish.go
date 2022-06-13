package service

import (
	"errors"
	"github.com/kuanyuh/simple-tiktok/dao"
	"github.com/kuanyuh/simple-tiktok/qiniu"
	"io"
	"time"
)

func PublishList(userId string) []Video{
	var videos []Video
	dao.DB.Table("video").Where("author_id = ?", userId).Order("created_at DESC").Find(&videos)
	return videos
}

func Publish(authorId int64, file io.Reader, fileSize int64, key string) (int, error) {
	//code, url := common.UploadLocal(key)
	code, url := qiniu.Upload(file, fileSize, key)
	if code == 200 {
		video := Video{
			AuthorId: authorId,
			PlayUrl: url,
			CoverUrl: "http://rdaxet1xt.hn-bkt.clouddn.com/test%2Fdark.png",
			FavoriteCount: 0,
			CommentCount: 0,
			Title: "",
			CreatedAt: time.Now().Unix()}
		dao.DB.Table("video").Create(video)
		return 200, nil
	}
	return 100, errors.New("文件上传失败")
}