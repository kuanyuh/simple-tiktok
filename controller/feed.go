package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kuanyuh/simple-tiktok/service"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	videos := service.Feed()
	feedVideos := []Video{}
	//结构体复制
	for _, video := range videos {
		feedVideo := Video{
			Id: video.Id,
			Author: User(service.GetUserinfoById(string(video.AuthorId))),
			PlayUrl: video.PlayUrl,
			CoverUrl: video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount: video.CommentCount,
			IsFavorite: video.IsFavorite,
		}
		feedVideos = append(feedVideos, feedVideo)
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: feedVideos,
		NextTime:  time.Now().Unix(),
	})
}

func videoCopy(feedvideos *[]Video, videos *[]service.Video)  {

}