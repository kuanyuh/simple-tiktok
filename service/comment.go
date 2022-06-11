package service

import (
	"github.com/kuanyuh/simple-tiktok/dao"
	"time"
)

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	VideoId    int64  `json:"video_id,omitempty"`
	UserId     int64  `json:"user_id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

func SaveComment(videoId int64, userId int64, content string) Comment {
	comment := Comment{
		VideoId:    videoId,
		UserId:     userId,
		Content:    content,
		CreateDate: time.Time.Format(time.Now(), "01-02"),
	}
	dao.DB.Table("comment").Save(&comment)
	var count int64
	dao.DB.Table("comment").Where("video_id = ?", videoId).Count(&count)
	dao.DB.Table("video").Where("id = ?", videoId).Set("comment_count", count)
	return comment
}

func DeleteComment(commentId int64) {
	dao.DB.Table("comment").Where("id = ?", commentId).Update("deleted_at", time.Now().Unix())

}

func GetComments(videoId int64) []Comment {
	var comments []Comment
	dao.DB.Table("comment").Where("video_id = ? and deleted_at is null", videoId).Order("create_date desc").Find(&comments)
	return comments
}
