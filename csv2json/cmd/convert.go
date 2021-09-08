package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/blackhorseya/csv2json/internal/business/convert"
	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("file")
		if err != nil {
			fmt.Println(err)
			return
		}

		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Println(err)
			return
		}

		business := convert.NewImpl()
		users, err := business.LoadFromCSV(string(data))
		if err != nil {
			fmt.Println(err)
			return
		}

		bytes, err := json.Marshal(users)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(bytes))
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	convertCmd.Flags().StringP("file", "f", "./test/data/raw.csv", "specific source file path")
}
