package service

import (
	"fmt"
	"github.com/kuanyuh/simple-tiktok/dao"
)

type Relation struct {
	Id       int64 `json:"id,omitempty"`
	UserId   int64 `json:"user_id,omitempty"`
	ToUserId int64 `json:"to_user_id,omitempty"`
	IsFollow bool  `json:"is_follow,omitempty"`
}

//DoRelation 关注功能
func DoRelation(userId int64, toUserId int64) {
	newRelation := Relation{UserId: userId, ToUserId: toUserId, IsFollow: true}
	dao.DB.Table("relation").Save(&newRelation)
}

//DelRelation 取消关注
func DelRelation(userId int64, toUserId int64) {
	r := GetRelation(userId, toUserId)
	dao.DB.Table("relation").Delete(&r)
}

func GetRelation(userId int64, toUserId int64) Relation {
	relation := Relation{}
	dao.DB.Table("relation").Where("user_id = ? AND to_user_id = ?", userId, toUserId).Find(&relation)
	return relation
}

//FollowList 关注列表
func FollowList(userId int64) []User {
	var users []User
	var relations []Relation
	dao.DB.Table("relation").Where("user_id = ?", userId).Find(&relations)
	for _, relation := range relations {
		user := GetUserinfoById(fmt.Sprint(relation.ToUserId))
		users = append(users, user)
	}
	return users
}

//FollowerList 粉丝列表
func FollowerList(userId int64) []User {
	var users []User
	var relations []Relation
	dao.DB.Table("relation").Where("to_user_id = ?", userId).Find(&relations)
	for _, relation := range relations {
		user := GetUserinfoById(fmt.Sprint(relation.UserId))
		users = append(users, user)
	}
	return users
}
