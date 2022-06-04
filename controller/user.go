package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kuanyuh/simple-tiktok/common"
	"github.com/kuanyuh/simple-tiktok/service"
	"net/http"
	"strconv"
)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

//Register 用户注册
func Register(c *gin.Context) {
	//从请求参数获取登录名和密码
	username := c.Query("username")
	password := c.Query("password")

	if service.IsExist(username) {
		//如果邮箱被注册过了
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		//保存新用户，获取新注册用户id
		id := service.SaveUser(username, password)
		user := service.GetUserinfoById(strconv.FormatInt(id, 10))
		//根据用户信息生成token
		token := common.GetHStoken("JWT", &user)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   id,
			Token:    token,
		})
	}

}

//Login 用户登录
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if service.IsExist(username) {
		user := service.GetUser(username, password)
		if user == (service.UserLogin{}) {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "密码错误"},
			})
		} else {
			userInfo := service.GetUserinfoById(strconv.FormatInt(user.ID, 10))
			//根据用户信息生成token
			token := common.GetHStoken("JWT", &userInfo)
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   user.ID,
				Token:    token,
			})
		}
	} else { //用户不存在
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "用户不存在，请先注册"},
		})
	}
}

//UserInfo 跳转到用户信息页面
func UserInfo(c *gin.Context) {
	token := c.Query("token")

	//解析token
	claims := common.ParseHStoken(token)
	id, _ := json.Marshal(claims["id"])

	user := service.GetUserinfoById(string(id))
	if user != (service.User{}) {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     User(user),
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "用户不存在"},
		})
	}

}
