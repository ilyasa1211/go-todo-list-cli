/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [title] [description]",
	Short: "Create a todo",
	Long:  `Create a todo`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.OpenFile("data/data.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

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

		w.Write([]string{title, description})
		w.Flush()

		fmt.Println("create called")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
