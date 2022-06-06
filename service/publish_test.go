package service

import (
	"fmt"
	"github.com/kuanyuh/simple-tiktok/dao"
	"testing"
)

func TestPublishList(t *testing.T) {
	dao.Init()
	for _, video := range PublishList(10) {
		fmt.Printf("%v\n",video)
	}
}
