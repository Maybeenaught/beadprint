package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(convertCmd)
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Ingests an image file (jpg/png) and converts the content to a printable beading template",
	Long:  "Ingests an image file (jpg/png) and converts the content to a printable beading template",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("not implemented yet")
	},
}
