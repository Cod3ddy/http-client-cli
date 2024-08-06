/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package methods

import (
	"github.com/spf13/cobra"
)

// router.goCmd represents the router.go command
var RouterCmd = &cobra.Command{
	Use:   "router",
	Short: "Entry point of all http commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func addSubcommands() {
	RouterCmd.AddCommand(getCmd)
	RouterCmd.AddCommand(postCmd)
	RouterCmd.AddCommand(putCmd)
	RouterCmd.AddCommand(deleteCmd)
}

func init() {
	addSubcommands()
}
