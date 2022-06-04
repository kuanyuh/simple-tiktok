package controller

import (
	"fmt"
	"github.com/kuanyuh/simple-tiktok/dao"
	"github.com/kuanyuh/simple-tiktok/service"
	"testing"
)

func TestVideoCopy(t *testing.T) {
	dao.Init()
	videos:= service.Feed()
	for i, video := range videos {
		fmt.Printf(" %v: %v\n",i,video)
	}
	feedVideos := []Video{}
	videoCopy(&feedVideos,&videos)
	for _, video := range feedVideos {
		fmt.Printf("%v\n", video)
	}
}