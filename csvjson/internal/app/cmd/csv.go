package cmd

import (
	"fmt"
	"glasnostic/internal/app/biz/trans"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var (
	format string

	csvCmd = &cobra.Command{
		Use:   "csv",
		Short: "transfer json file to csv",
		Run: func(cmd *cobra.Command, args []string) {
			biz := trans.NewImpl()
			people, err := biz.LoadJson(file)
			if err != nil {
				fmt.Println(err)
				return
			}

			ret, err := biz.ToCSVByHeader(people, format)
			if err != nil {
				fmt.Println(err)
				return
			}

			if len(out) != 0 {
				// todo: 2021-01-30|13:47|doggy|improve it, overwrite file content
				err := ioutil.WriteFile(out, []byte(ret), 0644)
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			fmt.Println(ret)
		},
	}
)

func init() {
	rootCmd.AddCommand(csvCmd)

	csvCmd.Flags().StringVar(&format, "format", "FirstName,LastName,Email,Description,Role,Phone,ID", "output format")
}
