package main

import (
	"errors"
	"fmt"
	"github.com/cnbattle/upcloud/config"
	"github.com/cnbattle/upcloud/core/cloud"
	"github.com/cnbattle/upcloud/core/utils"
	"strings"
)

func main() {
	platform := config.GetEnv("UP_CLOUD_PLATFORM")
	commInterface, err := selectInterFace(platform)
	if err != nil {
		fmt.Println("selectInterFace error:", err)
		return
	}

	err = commInterface.Init()
	if err != nil {
		fmt.Println("commInterface.Init error:", err)
		return
	}
	// 获取已存在文件列表
	fmt.Println("Start to clear data.")
	getAll, err := commInterface.GetAll()
	if len(getAll) > 0 {
		// 删除已存在文件
		_ = commInterface.DelAll(getAll)
	}
	// 上传新的文件
	fmt.Println("Start uploading data.")
	files := utils.Local(config.GetEnv("UP_CLOUD_PATH"))
	for _, file := range files {
		fmt.Print(".")
		file = strings.Replace(file, "\\", "/", -1)
		err := commInterface.Upload(config.GetEnv("UP_CLOUD_PATH")+file, file)
		if err != nil {
			fmt.Println("commInterface.Upload error:", err)
			return
		}
	}
	fmt.Println()
	fmt.Println("Start refreshing CDN data.")
	prefetch := config.GetEnv("UP_CLOUD_PREFETCH_URLS")
	err = commInterface.Prefetch(strings.Split(prefetch, ","))
	if err != nil {
		fmt.Println("commInterface.Prefetch error:", err)
		return
	}
	fmt.Print("Successful !!!")
}

func selectInterFace(platform string) (cloud.CommInterface, error) {
	switch platform {
	case "qiniu":
		platformStruct := cloud.Qiniu{
			AccessKey: config.GetEnv("UP_CLOUD_ACCESS_KEY"),
			SecretKey: config.GetEnv("UP_CLOUD_SECRET_KEY"),
			Bucket:    config.GetEnv("UP_CLOUD_BUCKET"),
		}
		return &platformStruct, nil
	case "tencent":
		platformStruct := cloud.Tencent{
			SecretID:  config.GetEnv("UP_CLOUD_SECRET_ID"),
			SecretKey: config.GetEnv("UP_CLOUD_SECRET_KEY"),
			VisitNode: config.GetEnv("UP_CLOUD_VISIT_NODE"),
		}
		return &platformStruct, nil
	default:
		return nil, errors.New("config platform is error:" + platform)
	}
}
