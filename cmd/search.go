package cmd

import (
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search registry for modules",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: do the search
	},
}
