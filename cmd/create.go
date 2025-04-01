/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [title] [description]",
	Short: "Create a todo",
	Long:  `Create a todo`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.OpenFile("data/data.csv", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)

		if err != nil {
			log.Fatalln("Error occured when opening file: ", err)
		}

		defer f.Close()

		var title string = args[0]
		var description string

		if len(args) > 1 {
			description = args[1]
		}

		w := csv.NewWriter(f)
		r := csv.NewReader(f)

		records, err := r.ReadAll()
		var lastId int

		if err != nil {
			log.Fatal(err)
		}

		if len(records) > 0 {
			if lastId, err = strconv.Atoi(records[len(records)-1][0]); err != nil {
				log.Fatal(err)
			}
		}

		w.Write([]string{strconv.Itoa(lastId + 1), title, description})
		w.Flush()

		fmt.Println("create called")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
