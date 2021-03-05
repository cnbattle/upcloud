package main

import (
	"fmt"
	"strings"
	"sync"

	"github.com/cnbattle/upcloud/config"
	"github.com/cnbattle/upcloud/utils"

	"github.com/panjf2000/ants/v2"
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
	pool, _ := ants.NewPool(config.GetDefaultEnvToInt("UP_CLOUD_POOL_SIZE", 10))
	defer pool.Release()
	var wg sync.WaitGroup
	for _, file := range files {
		wg.Add(1)
		local := file.Local
		upKey := file.UpKey
		err = pool.Submit(func() {
			err := commInterface.Upload(local, upKey)
			if err != nil {
				fmt.Println("commInterface.Upload error:", err)
			}
			fmt.Print(".")
			wg.Done()
		})
		if err != nil {
			fmt.Println("pool.Submit error:", err)
		}
	}
	wg.Wait()
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
