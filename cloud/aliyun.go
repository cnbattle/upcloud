// Package cloud 支持的云平台
package cloud

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/cdn"
)

// Aliyun oss
type Aliyun struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	Bucket          string
	bucketClient    *oss.Bucket
}

// Init 初始化
func (t *Aliyun) Init() error {
	client, err := oss.New(t.Endpoint, t.AccessKeyID, t.AccessKeySecret)
	if err != nil {
		return err
	}
	t.bucketClient, err = client.Bucket(t.Bucket)
	return err
}

// GetAll 获取全部文件key
func (t *Aliyun) GetAll() (list []string, err error) {
	marker := ""
	for {
		lsRes, err := t.bucketClient.ListObjects(oss.Marker(marker))
		if err != nil {
			return nil, err
		}
		// 打印列举文件，默认情况下一次返回100条记录。
		for _, object := range lsRes.Objects {
			list = append(list, object.Key)
		}
		if lsRes.IsTruncated {
			marker = lsRes.NextMarker
		} else {
			break
		}
	}
	return
}

// DelAll 批量删除
func (t *Aliyun) DelAll(list []string) error {
	// 不返回删除的结果。
	_, err := t.bucketClient.DeleteObjects(list, oss.DeleteObjectsQuiet(true))
	return err
}

// Upload 上传
func (t *Aliyun) Upload(localFile, upKey string) error {
	err := t.bucketClient.PutObjectFromFile(upKey, localFile)
	return err
}

// Prefetch 刷新
func (t *Aliyun) Prefetch(urls []string) error {
	client, err := cdn.NewClientWithAccessKey("", t.AccessKeyID, t.AccessKeySecret)
	if err != nil {
		return err
	}
	request := cdn.CreateRefreshObjectCachesRequest()
	request.Scheme = "https"
	request.ObjectPath = strings.Join(urls, "\n")
	_, err = client.RefreshObjectCaches(request)
	return err
}
