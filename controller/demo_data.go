package controller

import "github.com/kuanyuh/simple-tiktok/service"

const HOST = "http://100.120.62.90:8080"

var DemoVideos = []Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            2,
		Author:        DemoUser,
		PlayUrl:       "https://oss-p56.xpccdn.com/prod/footage/preview/j2MvZd1p5G.mp4",
		CoverUrl:      "https://kuanyuh.github.io/img/favicon.png",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            3,
		Author:        DemoUser,
		PlayUrl:       HOST + "/static/v2.mp4",
		CoverUrl:      "https://kuanyuh.github.io/img/favicon.png",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            4,
		Author:        DemoUser,
		PlayUrl:       "http://121.5.225.209:9000/product/product_1654357481775.mp4",
		CoverUrl:      "http://121.5.225.209:9000/product/product_1654357233945.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            5,
		Author:        DemoUser,
		PlayUrl:       "http://121.5.225.209:9000/product/product_1654358443093.mp4",
		CoverUrl:      "http://121.5.225.209:9000/product/product_1654358548050.png",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []service.Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = service.User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
