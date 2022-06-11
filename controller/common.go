package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kuanyuh/simple-tiktok/common"
	"github.com/kuanyuh/simple-tiktok/service"
	"strconv"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

func CommentsFormat(userComments []service.Comment) []Comment {
	var comments []Comment
	for _, userComment := range userComments {
		comments = append(comments, CommentFormat(userComment))
	}
	return comments
}

func CommentFormat(commentModel service.Comment) Comment {
	var comment Comment
	comment = Comment{
		Id:         commentModel.Id,
		User:       User(service.GetUserinfoById(strconv.FormatInt(commentModel.UserId, 10))),
		Content:    commentModel.Content,
		CreateDate: commentModel.CreateDate,
	}
	return comment
}

func GetUserFromToken(c *gin.Context) service.User {
	token := c.Query("token")
	claims := common.ParseHStoken(token)
	id, _ := json.Marshal(claims["id"])
	user := service.GetUserinfoById(string(id))
	return user
}
