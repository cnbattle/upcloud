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
			var index int
			for i, i2 := range cloud.Platform {
				i++
				fmt.Printf("%v %s\n", i, i2)
			}
			fmt.Print("请选择平台编号：")
			fmt.Scanln(&index)
			fmt.Println("index:", index)
			fmt.Println("index data:", cloud.Platform[index-1])
		},
	}

}
