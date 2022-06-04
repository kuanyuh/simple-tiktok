package service

import (
	"fmt"
	"github.com/kuanyuh/simple-tiktok/dao"
	"testing"
)

func TestFeed(t *testing.T) {

	dao.Init()
	videos:= Feed()
	for i, video := range videos {
		fmt.Printf(" %v: %v",i,video)
	}
}
