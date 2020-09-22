package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/cnbattle/upcloud/config"
	"github.com/cnbattle/upcloud/core/cloud"
	"github.com/cnbattle/upcloud/core/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
)

var Create *cobra.Command

func init() {
	Create = &cobra.Command{
		Use:   "create",
		Short: "create deploy config data.",
		Long:  `create deploy config data.`,
		Run: func(cmd *cobra.Command, args []string) {
			// 选择平台
			var index int
			if len(cloud.Platform) > 0 {
				for i, i2 := range cloud.Platform {
					i++
					fmt.Printf("%v-%s\n", i, i2)
				}
				fmt.Print("Select Platform No：")
				fmt.Scanln(&index)
			} else {
				index = 1
			}
			platform, err := cloud.SelectPlatform(cloud.Platform[index-1])
			if err != nil {
				panic(err)
			}
			setting := platform.Setting()
			config.Conf = append(config.Conf, setting)
			b, err := json.Marshal(config.Conf)
			if err != nil {
				panic(err)
			}
			//生成json文件
			err = ioutil.WriteFile(utils.GetConfig(), b, 0666)
			if err != nil {
				panic(err)
			}
		},
	}

}
