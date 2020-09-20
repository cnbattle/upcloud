package cmd

import (
	"fmt"
	"github.com/cnbattle/upcloud/core/cloud"
	"github.com/spf13/cobra"
)

var Create *cobra.Command

func init() {
	Create = &cobra.Command{
		Use:   "create",
		Short: "create a config data.",
		Long:  `create a config data.`,
		Run: func(cmd *cobra.Command, args []string) {
			// 选择平台
			for i, i2 := range cloud.Platform {
				i++
				fmt.Printf("%v-%s\n", i, i2)
			}
			fmt.Print("请选择平台编号：")
			var index int
			fmt.Scanln(&index)
			platform, err := cloud.SelectPlatform(cloud.Platform[index-1])
			if err != nil {
				panic(err)
			}
			err = platform.Setting()
			if err != nil {
				panic(err)
			}
		},
	}

}
