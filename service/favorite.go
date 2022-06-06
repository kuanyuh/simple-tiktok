package service

import (
	"github.com/kuanyuh/simple-tiktok/dao"
)

type Favorite struct {
	Id         int64  `json:"id,omitempty"`
	UserId     string `json:"user_id"`
	VideoId    string `json:"video_id"`
	IsFavorite int    `json:"is_favorite,omitempty"` //标识符，1为点赞，2为取消点赞
}

func SaveFavorite(userId string, videoId string, actionType int) {
	if !IsExistFavorite(userId, videoId) {
		favorite := Favorite{UserId: userId, VideoId: videoId, IsFavorite: 1} //需要保存时必然是点赞操作
		dao.DB.Table("favorite").Save(&favorite)
	} else {
		ChangeFavorite(userId, videoId, actionType)
	}
}

// ChangeFavorite 取消操作，用户取消点赞后将is_favorite标识符改2
func ChangeFavorite(userId string, videoId string, actionType int) {
	var favorite Favorite
	dbRes := dao.DB.Table("favorite").Where("user_id = ? AND video_id = ?", userId, videoId).Find(&favorite)
	favorite.IsFavorite = actionType
	dbRes.Save(&favorite)
}

func UpdateVideo(videoId string, actionType int) {
	var video Video
	if actionType == 1 {
		dbRes := dao.DB.Table("video").Where("id = ?", videoId).Find(&video)
		video.FavoriteCount++
		dbRes.Save(&video)
	} else if actionType == 2 {
		dbRes := dao.DB.Table("video").Where("id = ?", videoId).Find(&video)
		video.FavoriteCount--
		dbRes.Save(&video)
	}
}

// GetFavoriteVideoList 获取用户点赞的视频列表
func GetFavoriteVideoList(userId string) []Video {
	var favorites []Favorite //在favorite表获取视频id列表
	var videos []Video
	dao.DB.Table("favorite").Where("user_id = ? AND is_favorite = ?", userId, "1").Find(&favorites)
	for i := 0; i < len(favorites); i++ {
		videos = append(videos, GetVideoById(favorites[i].VideoId))
	}
	return videos
}

func GetVideoById(videoId string) Video {
	var video Video
	dao.DB.Table("video").Where("id = ?", videoId).First(&video)
	return video
}

// IsExistFavorite @Description: 通过user_id和video_id查找点赞记录是否存在
//  @return bool：若存在返回true,此时不需要新增数据记录，将is_favorite改为1即可
//
func IsExistFavorite(userId string, videoId string) bool {
	favorite := Favorite{}
	dao.DB.Table("favorite").Where("user_id = ? AND video_id = ?", userId, videoId).First(&favorite)
	if favorite == (Favorite{}) { //不存在返回false
		return false
	}
	return true
}

func IsFavorite(userId string, videoId string) bool {
	favorite := Favorite{}
	dao.DB.Table("favorite").Where("user_id = ? AND video_id = ? AND is_favorite = ?", userId, videoId, "1").First(&favorite)
	if favorite == (Favorite{}) { //不存在返回false
		return false
	}
	return true
}
