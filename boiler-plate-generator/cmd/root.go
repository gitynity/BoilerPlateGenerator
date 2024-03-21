// cmd/root.go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "boiler-plate-generator",
	Short: "A generator for boilerplate code",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to boiler-plate-generator!")
		Generate()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
