/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Listing all your todos",
	Long:  `Retrieve all your todos in table view`,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.OpenFile("data/data.csv", os.O_RDONLY, 0644)

		if err != nil {
			log.Fatalln("Error occurred while opening file: ", err)
		}
		defer f.Close()

		w := tabwriter.NewWriter(os.Stdout, 4, 8, 4, '\t', tabwriter.AlignRight)

		r := csv.NewReader(f)

		records, err := r.ReadAll()

		if err != nil {
			log.Fatal(err)
		}

		for _, record := range records {
			fmt.Fprintln(w, record[0]+"\t"+record[1]+"\t")
		}

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
