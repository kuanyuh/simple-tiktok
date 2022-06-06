package controller

import (
	"fmt"
	"github.com/kuanyuh/simple-tiktok/dao"
	"github.com/kuanyuh/simple-tiktok/service"
	"testing"
)

func TestPublishList(t *testing.T) {
	dao.Init()
	videos := service.PublishList("10")
	var publishVideo []Video
	videoCopy(&publishVideo, &videos)

	fmt.Printf("%v",publishVideo)
}