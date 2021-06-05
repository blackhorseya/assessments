package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	file string
	out  string

	rootCmd = &cobra.Command{
		Use:  "glas",
		Long: "csvjson coding test",
	}
)

func init() {
	rootCmd.Version = "0.0.1"

	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "./test/raw.csv", "source data file path")
	rootCmd.PersistentFlags().StringVarP(&out, "output", "o", "", "output data file path")
}

// Execute a root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
