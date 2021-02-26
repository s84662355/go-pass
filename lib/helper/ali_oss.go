package helper

import (
	_ "fmt"
	//"strconv"
	_ "encoding/json"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	_ "strings"
	_ "time"
)

var ossEndpoint = ""
var ossAccessKeyID = ""
var ossAccessKeySecret = ""
var ossBucket = ""
var url = ""

type AliOss struct {
	ossCli *oss.Client
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t AliOss) Create() AliOss {
	tt := AliOss{}
	tt.ossCli, _ = oss.New(ossEndpoint, ossAccessKeyID, ossAccessKeySecret)
	return tt
}

// Bucket: 获取bucket存储空间
func (t *AliOss) bucket() *oss.Bucket {
	bucket, err := t.ossCli.Bucket(ossBucket)
	if err != nil {
		return nil
	}
	return bucket
}

func (t *AliOss) Upload(filename string, ossPath string) string {

	file, _ := os.Open(filename)

	err := t.bucket().PutObject(ossPath, file)
	if err != nil {

		return ""
	}
	return url + ossPath
}
