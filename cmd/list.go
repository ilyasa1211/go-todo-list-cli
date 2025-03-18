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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.OpenFile("data/data.csv", os.O_RDONLY, 0644)

		if err != nil {
			log.Fatalln("Error occurred while opening file: ", err)
		}
		defer f.Close()

		w := tabwriter.NewWriter(os.Stdout, 0, 8, 2, '\t', tabwriter.AlignRight)

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
