package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kuanyuh/simple-tiktok/common"
	"github.com/kuanyuh/simple-tiktok/service"
	"net/http"
	"path/filepath"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	//解析token
	claims := common.ParseHStoken(token)
	id, _ := json.Marshal(claims["id"])
	userInfo := service.GetUserinfoById(string(id))

	if userInfo == (service.User{}) {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", userInfo.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	//解析token
	claims := common.ParseHStoken(token)
	id, _ := json.Marshal(claims["id"])
	user := service.GetUserinfoById(string(id))
	if user == (service.User{}) {
		c.JSON(http.StatusOK, UserResponse{Response: Response{StatusCode: -1}})
	}
	userId := c.Query("user_id")
	videos := service.PublishList(userId)
	var publishVideo []Video
	videoCopy(&publishVideo, &videos, string(id))
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: publishVideo,
	})
}
