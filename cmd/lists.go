package cmd

import (
	"github.com/cnbattle/upcloud/config"
	"github.com/modood/table"
	"github.com/spf13/cobra"
)

// List 列表
var List *cobra.Command

func init() {
	List = &cobra.Command{
		Use:              "list",
		Short:            "look project lists.",
		TraverseChildren: true,
		Long:             `look project lists.`,
		Run: func(cmd *cobra.Command, args []string) {
			table.Output(config.Conf)
		},
	}

}
