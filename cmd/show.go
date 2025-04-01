/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show [id]",
	Short: "show only specific todo list by id",
	Long:  `show only specific todo list by id`,
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		filename := "data/data.csv"

		f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)

		if err != nil {
			log.Fatalln("error opening file: ", err)
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)

		w := tabwriter.NewWriter(os.Stdout, 4, 8, 4, '\t', tabwriter.AlignRight)
		for scanner.Scan() {
			row := scanner.Text()

			if strings.HasPrefix(row, id) {
				w.Write([]byte(row + "\n"))
				w.Flush()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
