package cmd

import (
	"errors"
	"fmt"
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
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a color argument")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			// 获取 args
			if !utils.IsExistProjectConfig(args[0]) {
				fmt.Println("project name is err")
				return
			}

			fmt.Println("success")
			// 获取已存在文件列表
			// 删除已存在文件
			// 上传新的文件
			// 刷新index.html
		},
	}

}
