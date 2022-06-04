package service

import (
	"github.com/kuanyuh/simple-tiktok/dao"
)


type Video struct {
	Id            int64  `json:"id,omitempty"`
	AuthorId        int64   `json:"author_id"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	LatestTime    int64  `json:"latest_time,omitempty"`
}

//Feed 视频流，app打开时
func Feed() []Video {
	var videos []Video
	dao.DB.Table("video").Find(&videos)
	return videos
}
