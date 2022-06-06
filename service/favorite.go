package service

import "github.com/kuanyuh/simple-tiktok/dao"

func AlreadyFollow(userId int64) {
	dao.DB.Table("user").Where("id = ?", userId).Update("is_follow", 1)
}

func NotFollow(userId int64) {
	dao.DB.Table("user").Where("id = ?", userId).Update("is_follow", 0)
}
