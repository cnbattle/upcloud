package main

import (
	"fmt"
	"github.com/cnbattle/upcloud/config"
	"github.com/cnbattle/upcloud/core/utils"
	"strings"
)

func main() {
	platform := config.GetEnvForPanic("UP_CLOUD_PLATFORM")
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
	files := utils.Local(config.GetEnvForPanic("UP_CLOUD_PATH"))
	for _, file := range files {
		fmt.Print(".")
		file = strings.Replace(file, "\\", "/", -1)
		err := commInterface.Upload(config.GetEnvForPanic("UP_CLOUD_PATH")+file, file)
		if err != nil {
			fmt.Println("commInterface.Upload error:", err)
			return
		}
	}
	fmt.Println()
	fmt.Println("Start refreshing CDN data.")
	prefetch := config.GetEnvForPanic("UP_CLOUD_PREFETCH_URLS")
	err = commInterface.Prefetch(strings.Split(prefetch, ","))
	if err != nil {
		fmt.Println("commInterface.Prefetch error:", err)
		return
	}
	fmt.Print("Successful !!!")
}
