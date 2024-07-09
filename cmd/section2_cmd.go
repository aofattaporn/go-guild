package cmd

import (
	"aofattaporn/golang-guild/internal/section2"
	"fmt"

	"github.com/spf13/cobra"
)

var section2Cmd = &cobra.Command{
	Use:   "section2",
	Short: "A brief description of Section2 on Golang Guild",
	Long:  "A brief description of Section2",
	RunE: func(cmd *cobra.Command, args []string) error {

		section2.Test = "Confirm Section"
		fmt.Println(section2.Test)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(section2Cmd)
}
