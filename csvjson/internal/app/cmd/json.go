package cmd

import (
	"encoding/json"
	"fmt"
	"glasnostic/internal/app/biz/trans"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var (
	jsonCmd = &cobra.Command{
		Use:   "json",
		Short: "transfer csv file to json",
		Run: func(cmd *cobra.Command, args []string) {
			pretty, err := cmd.Flags().GetBool("pretty")
			if err != nil {
				fmt.Println(err)
				return
			}

			num, err := cmd.Flags().GetInt("worker")
			if err != nil {
				fmt.Println(err)
				return
			}

			biz := trans.NewImpl()
			people, err := biz.LoadCSV(file, num)
			if err != nil {
				fmt.Println(err)
				return
			}

			var ret []byte
			if pretty {
				ret, err = json.MarshalIndent(people, "", "  ")
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				ret, err = json.Marshal(people)
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			if len(out) != 0 {
				// todo: 2021-01-29|16:22|doggy|improve it, overwrite file
				err := ioutil.WriteFile(out, ret, 0644)
				if err != nil {
					fmt.Println(err)
					return
				}
			}

			fmt.Println(string(ret))
		},
	}
)

func init() {
	rootCmd.AddCommand(jsonCmd)

	jsonCmd.Flags().Bool("pretty", false, "pretty json format")
	jsonCmd.Flags().IntP("worker", "w", 1, "configure worker numbers")
}
