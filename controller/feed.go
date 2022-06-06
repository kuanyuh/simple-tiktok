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
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	//解析token
	claims := common.ParseHStoken(token)
	id, _ := json.Marshal(claims["id"])
	user := service.GetUserinfoById(string(id))
	if user == (service.User{}) {
		c.JSON(http.StatusOK, UserResponse{Response: Response{StatusCode: -1}})
	}
	videos := service.Feed()
	feedVideos := []Video{}
	//结构体复制
	videoCopy(&feedVideos,&videos)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: feedVideos,
		NextTime:  time.Now().Unix(),
	})
}

func videoCopy(feedVideos *[]Video, videos *[]service.Video)  {
	for _, video := range *videos {
		feedVideo := Video{
			Id: video.Id,
			Author: User(service.GetUserinfoById(strconv.FormatInt(video.AuthorId,10))),
			PlayUrl: video.PlayUrl,
			CoverUrl: video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount: video.CommentCount,
			IsFavorite: video.IsFavorite,
		}
		*feedVideos = append(*feedVideos, feedVideo)
	}
}