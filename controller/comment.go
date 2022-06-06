package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kuanyuh/simple-tiktok/common"
	"github.com/kuanyuh/simple-tiktok/service"
	"net/http"
)

type CommentListResponse struct {
	service.Response
	CommentList []service.Comment `json:"comment_list,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")

	//解析token
	claims := common.ParseHStoken(token)
	id, _ := json.Marshal(claims["id"])

	userInfo := service.GetUserinfoById(string(id))

	if userInfo != (service.User{}) {
		c.JSON(http.StatusOK, service.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    service.Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
