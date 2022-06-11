package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kuanyuh/simple-tiktok/service"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	user := GetUserFromToken(c)
	actionType := c.Query("action_type")
	if user != (service.User{}) {
		if actionType == "1" {
			// 从 c 创建 comment
			videoId, err1 := strconv.ParseInt(c.Query("video_id"), 10, 64)
			if err1 != nil {
				return
			}
			// 存入数据库
			comment := service.SaveComment(videoId, user.Id, c.Query("comment_text"))

			// response
			c.JSON(http.StatusOK, CommentResponse{
				Response: Response{
					StatusCode: 0,
					StatusMsg:  "Comment succeed",
				},
				Comment: CommentFormat(comment),
			})
		} else {
			commentId, err3 := strconv.ParseInt(c.Query("comment_id"), 10, 64)
			if err3 != nil {
				return
			}
			// 从数据库删除
			service.DeleteComment(commentId)
			// 返回
			c.JSON(http.StatusOK, Response{
				StatusCode: 0,
				StatusMsg:  "comment has been deleted",
			})
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {

	user := GetUserFromToken(c)
	if user != (service.User{}) {
		videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
		if err != nil {
			return
		}
		comments := service.GetComments(videoId)
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    Response{StatusCode: 0},
			CommentList: CommentsFormat(comments),
		})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}
