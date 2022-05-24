package service

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kuanyuh/simple-tiktok/dao"
	"math/rand"
	"time"
)

type UserLogin struct {
	ID int64 //gorm会默认把ID当作主键
	//gorm.Model
	Email    string
	Password string
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

// IsExist
//  @Description: 通过注册邮箱查找用户是否存在
//  @param email: 注册邮箱号
//  @return bool：若存在返回true
//
func IsExist(email string) bool {
	user := UserLogin{}
	dao.DB.Table("login").Where("email = ?", email).First(&user)
	if user == (UserLogin{}) { //不存在返回false
		return false
	}
	return true
}

// SaveUser
//  @Description: 保存新用户到user表（先保存到login表再保存到user表，用代码模拟外键保证一致性）
//  @param username：注册时的邮箱号
//  @param password：注册时的密码
//  @return int64：用户的id
func SaveUser(username string, password string) int64 {
	user := UserLogin{Email: username, Password: password}
	dao.DB.Table("login").Save(&user)
	//获取新建用户的id
	user1 := GetUser(username, password)
	//将新建用户同时保存到user表中，代码模拟外键
	userInfo := User{Id: user1.ID, Name: GetRandomString(8), FollowCount: 0, FollowerCount: 0, IsFollow: false}
	dao.DB.Table("user").Save(&userInfo)

	return user1.ID
}

// GetUser
//  @Description: 获取UserLogin
//  @param username：注册邮箱号
//  @param password：注册密码
//  @return UserLogin：已注册用户的UserLogin对象
func GetUser(username string, password string) UserLogin {
	user := UserLogin{}
	dao.DB.Table("login").Where("email = ? AND password = ?", username, password).First(&user)
	return user
}

// GetUserinfoById
//  @Description: 通过id获取用户详细信息（User对象）
//  @param id：用户id
//  @return User：返回用户信息（User）
func GetUserinfoById(id string) User {
	user := User{}
	dao.DB.Table("user").Where("id = ?", id).First(&user)
	return user
}

//生成随机字符串，作为用户名
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
