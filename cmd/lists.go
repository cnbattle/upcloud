package cmd

import (
	"fmt"
	"github.com/cnbattle/upcloud/core/utils"
	"github.com/spf13/cobra"
	"strings"
)

var List *cobra.Command

func init() {
	List = &cobra.Command{
		Use:              "list",
		Short:            "look project lists.",
		TraverseChildren: true,
		Long:             `look project lists.`,
		Run: func(cmd *cobra.Command, args []string) {
			files := utils.Local(utils.GetConfigDir())
			for _, file := range files {
				split := strings.Split(file, ".")
				//fmt.Println(i,split[0])
				fmt.Printf("Project Name:%v\tConfig Path:%v\n", split[0], utils.GetExistProjectConfig(split[0]))
			}
		},
	}

}
