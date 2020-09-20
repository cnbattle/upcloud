package cmd

import (
	"fmt"
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
			//if !utils.IsExistProjectConfig(args[0]) {
			//	fmt.Println("project name is err")
			//	return
			//}
			// 读取配置

			fmt.Println("success")
			// 获取已存在文件列表
			// 删除已存在文件
			// 上传新的文件
			// 刷新index.html
		},
	}

}
