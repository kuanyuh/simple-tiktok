package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kuanyuh/simple-tiktok/common"
	"github.com/kuanyuh/simple-tiktok/service"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	service.Response
	UserList []service.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	//解析token
	claims := common.ParseHStoken(token)
	id, _ := json.Marshal(claims["id"])
	userInfo := service.GetUserinfoById(string(id))
	if userInfo != (service.User{}) {
		author := c.Query("to_user_id")
		authorId, _ := strconv.ParseInt(author, 10, 64)
		userId, _ := strconv.ParseInt(string(id), 10, 64)
		relation := service.GetRelation(userId, authorId)
		if (relation != service.Relation{}) {
			//取关
			service.DelRelation(userId, authorId)
		} else {
			//关注
			service.DoRelation(userId, authorId)
		}
		c.JSON(http.StatusOK, service.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, service.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: service.Response{
			StatusCode: 0,
		},
		UserList: []service.User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: service.Response{
			StatusCode: 0,
		},
		UserList: []service.User{DemoUser},
	})
}
