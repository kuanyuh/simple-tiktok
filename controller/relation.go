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
	Response
	UserList []User `json:"user_list"`
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
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
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
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err!=nil {
		return
	}
	var users []User
	for _, user := range service.FollowList(uid) {
		id, _ := strconv.ParseInt(string(id), 10, 64)
		service.IsFollow(id, &user)
		users = append(users, User(user))
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
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
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err!=nil {
		return
	}
	var users []User
	for _, user := range service.FollowerList(uid) {
		id, _ := strconv.ParseInt(string(id), 10, 64)
		service.IsFollow(id, &user)
		users = append(users, User(user))
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}
