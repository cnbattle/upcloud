package main

import (
	"github.com/cnbattle/upcloud/cmd"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "upcloud"}
	rootCmd.AddCommand(cmd.Create, cmd.Delete, cmd.List, cmd.Deploy)
	rootCmd.Execute()
}
