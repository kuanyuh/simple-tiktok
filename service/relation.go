package service

import "github.com/kuanyuh/simple-tiktok/dao"

type Relation struct {
	Id       int64 `json:"id,omitempty"`
	UserId   int64 `json:"user_id,omitempty"`
	ToUserId int64 `json:"to_user_id,omitempty"`
	IsFollow bool  `json:"is_follow,omitempty"`
}

// DoRelation 关注功能
func DoRelation(userId int64, toUserId int64) {
	newRelation := Relation{UserId: userId, ToUserId: toUserId, IsFollow: true}
	dao.DB.Table("relation").Save(&newRelation)
}

// DelRelation 取消关注
func DelRelation(userId int64, toUserId int64) {
	r := GetRelation(userId, toUserId)
	dao.DB.Table("relation").Delete(&r)
}

func GetRelation(userId int64, toUserId int64) Relation {
	relation := Relation{}
	dao.DB.Table("relation").Where("user_id = ? AND to_user_id = ?", userId, toUserId).Find(&relation)
	return relation
}

func AlreadyFollow(userId int64) {
	dao.DB.Table("user").Where("id = ?", userId).Update("is_follow", 1)
}

func NotFollow(userId int64) {
	dao.DB.Table("user").Where("id = ?", userId).Update("is_follow", 0)
}
