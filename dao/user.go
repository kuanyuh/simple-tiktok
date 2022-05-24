package dao

import (
	"fmt"
)

type User struct {
	Id            int64  `json:"id,omitempty"`
	Account       string `json:"account,omitempty"`
	Password      string `json:"password,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}


// GetUser 查询用户
func GetUser(account string) (User, int) {

	//查询用db.Raw,其他用db.Exec
	var user User
	err := DB.Raw("select * from user where account = ? limit 1", account).Scan(&user).Error
	if err!=nil{
		return user, 0
	}
	fmt.Println("查询结果-user:",user)
	return user, 1
}

// CreateUser 新增用户
func CreateUser(user *User) int {
	//user.Password = ScryptPw(user.Password)
	err := DB.Create(&user).Error
	if err != nil {
		return 0 // 500
	}
	return 1
}