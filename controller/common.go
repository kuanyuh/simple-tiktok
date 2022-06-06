package controller

import "github.com/kuanyuh/simple-tiktok/service"

type Video struct {
	Id            int64        `json:"id,omitempty"`
	Author        service.User `json:"author"`
	PlayUrl       string       `json:"play_url"`
	CoverUrl      string       `json:"cover_url"`
	FavoriteCount int64        `json:"favorite_count"`
	CommentCount  int64        `json:"comment_count"`
	IsFavorite    bool         `json:"is_favorite"`
	Title         string       `json:"title"`
}
