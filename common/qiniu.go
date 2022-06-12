package common

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
)

var cfg storage.Config

func QiniuInit() {
	cfg = storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
}

func Upload(file io.Reader, fileSize int64,key string)(int, string){
	bucket := "octaneky"
	accessKey := "LgaemkTYojopB-PubyLBBoGxf0iBMyDtL5oLmnEQ"
	secretKey := "vjqGf90uLLu9EDbmxXh8sDZZ0h9HI8PSN3jguQ6v"


	putPolicy := storage.PutPolicy{
		Scope:               bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	//err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	err := formUploader.Put(context.Background(), &ret, upToken, key, file, fileSize, &putExtra)
	if err != nil {
		fmt.Println(err)
		return 100, err.Error()
	}
	url := "http://rdaxet1xt.hn-bkt.clouddn.com" + ret.Key
	fmt.Println(ret.Key,ret.Hash)
	return 200, url
}


func UploadLocal(key string)(int, string){
	bucket := "octaneky"
	accessKey := "LgaemkTYojopB-PubyLBBoGxf0iBMyDtL5oLmnEQ"
	secretKey := "vjqGf90uLLu9EDbmxXh8sDZZ0h9HI8PSN3jguQ6v"
	localFile := "D:\\code\\go\\golang\\bytedance\\simple-tiktok\\public\\bear.mp4"

	putPolicy := storage.PutPolicy{
		Scope:               bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	//err := formUploader.Put(context.Background(), &ret, upToken, key, file, fileSize, &putExtra)
	if err != nil {
		fmt.Println(err)
		return 100, err.Error()
	}
	url := "http://rdaxet1xt.hn-bkt.clouddn.com" + ret.Key
	fmt.Println(ret.Key,ret.Hash)
	return 200, url
}

// 自定义返回值结构体
type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}
