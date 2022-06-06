package service

import "github.com/kuanyuh/simple-tiktok/dao"

func PublishList(userId string) []Video{
	var videos []Video
	dao.DB.Table("video").Where("author_id = ?", userId).Order("created_at DESC").Find(&videos)
	return videos
}
