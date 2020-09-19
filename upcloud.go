package main

import (
	"context"
	"fmt"
	"github.com/cnbattle/upcloud/config"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/cdn"
	"github.com/qiniu/api.v7/v7/sms/rpc"
	"github.com/qiniu/api.v7/v7/storage"
	"log"
)

var upToken = ""
var bucketManager *storage.BucketManager
var mac *qbox.Mac
var pwd = config.GetDefaultEnv("pwd", "dist")

func main() {
	// 获取已存在文件列表
	list, err := getAll()
	log.Println("list err:", err)
	if len(list) > 0 {
		// 删除已存在文件
		log.Println("del err:", delAll(list))
	}

	// 上传新的文件
	files := local(pwd)
	for _, file := range files {
		log.Println("up:", file)
		err := upload(pwd+"/"+file, file)
		if err != nil {
			panic(err)
		}
	}
	// 刷新index.html
	log.Println("prefetch:", prefetch())
}

func getAll() (list []string, err error) {
	limit := 1000
	prefix := ""
	delimiter := ""
	//初始列举marker为空
	marker := ""
	for {
		entries, _, nextMarker, hasNext, err := bucketManager.ListFiles(bucket, prefix, delimiter, marker, limit)
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

func delAll(keys []string) error {
	deleteOps := make([]string, 0, len(keys))
	for _, key := range keys {
		deleteOps = append(deleteOps, storage.URIDelete(bucket, key))
	}
	rets, err := bucketManager.Batch(deleteOps)
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

func prefetch() error {
	cdnManager := cdn.NewCdnManager(mac)
	//刷新链接，单次请求链接不可以超过10个，如果超过，请分批发送请求
	urlsToRefresh := []string{
		"http://h5.ygxsj.com/",
	}
	_, err := cdnManager.RefreshDirs(urlsToRefresh)
	return err
}

func upload(localFile, key string) error {
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, nil)
	if err != nil {
		return err
	}
	return nil
}
