package service

import (
	"fmt"
	"github.com/kuanyuh/simple-tiktok/dao"
	"testing"
)

func TestFollowList(t *testing.T) {
	dao.Init()
	for _, user := range FollowList(1) {
		IsFollow(1, &user)
		fmt.Printf("%v\n", user)
	}
}

func TestFollowerList(t *testing.T) {
	dao.Init()
	for _, user := range FollowerList(1) {
		IsFollow(1, &user)
		fmt.Printf("%v\n", user)
	}
}