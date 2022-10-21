// Package cloud 支持的云平台
package cloud

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/cdn"
	"github.com/qiniu/go-sdk/v7/sms/rpc"
	"github.com/qiniu/go-sdk/v7/storage"
)

// Qiniu 七牛云
type Qiniu struct {
	AccessKey     string                 `json:"access_key"`
	SecretKey     string                 `json:"secret_key"`
	Bucket        string                 `json:"bucket"`
	UpToken       string                 `json:"-"`
	Mac           *qbox.Mac              `json:"-"`
	BucketManager *storage.BucketManager `json:"-"`
}

// Init 初始化
func (q *Qiniu) Init() error {
	putPolicy := storage.PutPolicy{
		Scope: q.Bucket,
	}
	q.Mac = qbox.NewMac(q.AccessKey, q.SecretKey)

	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	q.BucketManager = storage.NewBucketManager(q.Mac, &cfg)

	q.UpToken = putPolicy.UploadToken(q.Mac)
	return nil
}

// GetAll 获取全部文件key
func (q *Qiniu) GetAll() (list []string, err error) {
	limit := 1000
	prefix := ""
	delimiter := ""
	//初始列举marker为空
	marker := ""
	for {
		entries, _, nextMarker, hasNext, err := q.BucketManager.ListFiles(q.Bucket, prefix, delimiter, marker, limit)
		if err != nil {
			return nil, err
		}
		//print entries
		for _, entry := range entries {
			list = append(list, entry.Key)
		}
		if hasNext {
			marker = nextMarker
		} else {
			//list end
			break
		}
	}
	return list, nil
}

// DelAll 批量删除
func (q *Qiniu) DelAll(list []string) error {
	deleteOps := make([]string, 0, len(list))
	for _, key := range list {
		deleteOps = append(deleteOps, storage.URIDelete(q.Bucket, key))
		if len(deleteOps) == 900 {
			rets, err := q.BucketManager.Batch(deleteOps)
			if err != nil {
				// 遇到错误
				if _, ok := err.(*rpc.ErrorInfo); ok {
					for _, ret := range rets {
						// 200 为成功
						fmt.Printf("%d\n", ret.Code)
						if ret.Code != 200 {
							fmt.Printf("%s\n", ret.Data.Error)
						}
					}
				} else {
					fmt.Printf("batch error, %s", err)
				}
			}
			deleteOps = make([]string, 0, len(list))
		}
	}
	rets, err := q.BucketManager.Batch(deleteOps)
	if err != nil {
		// 遇到错误
		if _, ok := err.(*rpc.ErrorInfo); ok {
			for _, ret := range rets {
				// 200 为成功
				fmt.Printf("%d\n", ret.Code)
				if ret.Code != 200 {
					fmt.Printf("%s\n", ret.Data.Error)
				}
			}
		} else {
			fmt.Printf("batch error, %s", err)
		}
	}
	return nil
}

// Upload 上传
func (q *Qiniu) Upload(localFile, upKey string) error {
	err := q.Init()
	if err != nil {
		return err
	}
	cfg := storage.Config{}
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err = formUploader.PutFile(context.Background(), &ret, q.UpToken, upKey, localFile, nil)
	if err != nil {
		return err
	}
	return nil
}

// Prefetch 刷新
func (q *Qiniu) Prefetch(urls []string) error {
	cdnManager := cdn.NewCdnManager(q.Mac)
	_, err := cdnManager.RefreshUrls(urls)
	_, err = cdnManager.PrefetchUrls(urls)
	return err
}
