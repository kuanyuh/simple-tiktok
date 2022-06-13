package service

import (
	"fmt"
	"github.com/kuanyuh/simple-tiktok/dao"
	"github.com/kuanyuh/simple-tiktok/qiniu"
	"testing"
)

func TestPublishList(t *testing.T) {
	dao.Init()
	for _, video := range PublishList("10") {
		fmt.Printf("%v\n",video)
	}
}

func TestPublish(t *testing.T) {
	dao.Init()
	qiniu.QiniuInit()

}