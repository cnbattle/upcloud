package main

import (
	"github.com/cnbattle/upcloud/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "upcloud"}
	rootCmd.AddCommand(cmd.Create)
	rootCmd.Execute()
	// 获取已存在文件列表
	// 删除已存在文件
	// 上传新的文件
	// 刷新index.html
}
