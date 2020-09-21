package cmd

import (
	"errors"
	"fmt"
	"github.com/cnbattle/upcloud/config"
	"github.com/cnbattle/upcloud/core/cloud"
	"github.com/cnbattle/upcloud/core/utils"
	"github.com/spf13/cobra"
)

var Deploy *cobra.Command

func init() {
	Deploy = &cobra.Command{
		Use:              "deploy [PROJECT NAME]",
		Short:            "deploy a config",
		TraverseChildren: true,
		Long:             `deploy a config to server.`,
		Run: func(cmd *cobra.Command, args []string) {
			// 获取 args

			for _, name := range args {

				projectConfig, err := config.GetProjectConfig(name)
				if err != nil {
					fmt.Println(name, "deploy error:", err)
					continue
				}
				commInterface, err := selectInterFace(projectConfig)
				if err != nil {
					fmt.Println("selectInterFace error:", err)
					continue
				}

				err = commInterface.Init()
				if err != nil {
					fmt.Println("commInterface.Init error:", err)
					continue
				}
				// 获取已存在文件列表
				getAll, err := commInterface.GetAll()
				if len(getAll) > 0 {
					// 删除已存在文件
					_ = commInterface.DelAll(getAll)
				}

				// 上传新的文件
				files := utils.Local(projectConfig.Path)
				for _, file := range files {
					err := commInterface.Upload(projectConfig.Path+file, file)
					if err != nil {
						fmt.Println("commInterface.Upload error:", err)
					}
				}
				// 刷新index.html
				err = commInterface.Prefetch()
				if err != nil {
					fmt.Println("commInterface.Prefetch error:", err)
					continue
				}

			}
			fmt.Println("success")
		},
	}

}

func selectInterFace(conf config.ProjectConfig) (cloud.CommInterface, error) {
	switch conf.Platform {
	case "qiniu":
		qiniu := cloud.Qiniu{
			AccessKey: conf.Args["accessKey"],
			SecretKey: conf.Args["secretKey"],
			Bucket:    conf.Args["bucket"],
		}
		return &qiniu, nil
	default:
		return nil, errors.New("config platform is error:" + conf.Platform)
	}
}
