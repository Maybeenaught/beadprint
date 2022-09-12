package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "beadprint",
	Short: "Beadprint is a set of utilities designed to assist with crafting beadwork",
	Long:  "Beadprint is a set of utilities designed to assist with crafting beadwork",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello world")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
