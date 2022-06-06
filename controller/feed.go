package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kuanyuh/simple-tiktok/common"
	"github.com/kuanyuh/simple-tiktok/service"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	service.Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	videos := service.Feed()
	var feedVideos []Video
	token := c.Query("token")
	var userId string
	//结构体复制
	if token != "" { //如果有token(即用户登录了)
		claims := common.ParseHStoken(token)
		id, _ := json.Marshal(claims["id"])
		userId = string(id)
	}
	VideoCopy(&feedVideos, &videos, userId)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  service.Response{StatusCode: 0},
		VideoList: feedVideos,
		NextTime:  time.Now().Unix(),
	})
}

func VideoCopy(feedVideos *[]Video, videos *[]service.Video, userId string) {
	for _, video := range *videos {
		var isFavorite bool
		if userId == "" {
			isFavorite = false
		} else {
			isFavorite = service.IsFavorite(userId, strconv.FormatInt(video.Id, 10))
		}
		feedVideo := Video{
			Id:            video.Id,
			Author:        service.GetUserinfoById(strconv.FormatInt(video.AuthorId, 10)),
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		}

		*feedVideos = append(*feedVideos, feedVideo)
	}
}
