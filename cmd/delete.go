package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/cnbattle/upcloud/config"
	"github.com/cnbattle/upcloud/core/utils"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strings"
)

// Delete 删除命令
var Delete *cobra.Command

func init() {
	Delete = &cobra.Command{
		Use:   "delete",
		Short: "delete a config data.",
		Long:  `delete a config data.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("need some ProjectName")
				return
			}
			for _, projectName := range args {
				fmt.Printf("ProjectName:%v,", projectName)
				var newConf []config.ProjectConfig
				var has bool
				for _, project := range config.Conf {
					if strings.EqualFold(projectName, project.ProjectName) {
						has = true
					} else {
						newConf = append(newConf, project)
					}
				}
				// 不存在
				if !has {
					fmt.Println("ProjectName does not exist:" + projectName)
					continue
				}
				config.Conf = newConf
				b, err := json.Marshal(newConf)
				if err != nil {
					fmt.Println("json.Marshal error:" + err.Error())
					continue
				}
				//生成json文件
				err = ioutil.WriteFile(utils.GetConfig(), b, 0666)
				if err != nil {
					fmt.Println("ioutil.WriteFile error:" + err.Error())
					continue
				}
				fmt.Println("delete success.")
			}
		},
	}

}
