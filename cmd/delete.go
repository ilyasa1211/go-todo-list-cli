/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a todo by ID",
	Long:  `Delete a todo by ID`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
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

		fmt.Println("delete called")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
