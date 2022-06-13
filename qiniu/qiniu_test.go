package qiniu

import (
	"fmt"
	"testing"
)

func TestUploadLocal(t *testing.T) {

	QiniuInit()
	code, url := UploadLocal("test/bear.mp4")
	if code==200 {
		fmt.Printf("%v", url)
	}else if code==100 {
		fmt.Printf("%v", code)
	}
}
