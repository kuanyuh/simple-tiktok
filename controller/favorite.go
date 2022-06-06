package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kuanyuh/simple-tiktok/common"
	"github.com/kuanyuh/simple-tiktok/service"
	"net/http"
	"strconv"
)

// FavoriteAction 用户点赞操作
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	//解析token
	claims := common.ParseHStoken(token)
	id, _ := json.Marshal(claims["id"])
	userInfo := service.GetUserinfoById(string(id))
	if userInfo != (service.User{}) {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
	//点赞操作  分两步： 1.搜索到对应的视频点赞总数加1,2.添加数据到数据库favorite表，用于拉取用户点赞视频列表
	videoId := c.Query("video_id")
	actionType := c.Query("action_type") //1表示点赞，2表示取消点赞
	aType, _ := strconv.Atoi(actionType) //字符串转整形
	service.SaveFavorite(string(id), videoId, aType)
	service.UpdateVideo(videoId, aType)
}

func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	claims := common.ParseHStoken(token)
	//token id
	id, _ := json.Marshal(claims["id"])
	//验证 token 是否正确
	userInfo := service.GetUserinfoById(string(id))
	if userInfo == (service.User{}) {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{StatusCode: -1, StatusMsg: "incorrect token"},
		})
	}
	//验证前端传来的 user_id 是否正确
	userId := c.Query("user_id")
	userInfo = service.GetUserinfoById(userId)
	if userInfo == (service.User{}) {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{StatusCode: -1, StatusMsg: "incorrect token"},
		})
	}
	//通过用户id获取点赞视频列表
	videos := service.GetFavoriteVideoList(userId)
	var resVideos []Video
	videoCopy(&resVideos, &videos, userId)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: resVideos,
	})

}
