package service

import (
	"github.com/kuanyuh/simple-tiktok/dao"
)

//Feed 视频流，app打开时
func Feed() []Video {
	var videos []Video
	dao.DB.Table("video").Order("created_at DESC").Find(&videos)
	return videos
}
