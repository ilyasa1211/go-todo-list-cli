/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [id] [title] [description]",
	Short: "Update a todo by ID",
	Long:  `Update a todo by ID`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		title := args[1]
		desc := ""
		hasDesc := len(args) > 2

		if hasDesc {
			desc = args[2]
		}

		filename := "data/data.csv"

		f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)

		if err != nil {
			log.Fatalln("error opening file: ", err)
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)

		var newData [][]string

		for scanner.Scan() {
			row := scanner.Text()

			if strings.HasPrefix(row, id) {
				oldData := strings.Split(row, ",")
				if !hasDesc {
					desc = oldData[2]
				}
				newData = append(newData, []string{oldData[0], title, desc})
				continue
			}

			newData = append(newData, strings.Split(row, ","))
		}

		f, err = os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, 0644)

		if err != nil {
			log.Fatalln("error opening file", err)
		}

		defer f.Close()

		w := csv.NewWriter(f)

		w.WriteAll(newData)
		w.Flush()

	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
