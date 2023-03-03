package cmd

import (
	"github.com/TheMorgz/gobirth/internal/input"
	"github.com/spf13/cobra"
)

var reader input.Reader

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Reads a static input file and outputs results of today's birthdays.",
	Long:  `This command will read our static file(located under static > files > input.json). It will then loop through the records and analyse birthday data.`,
	Run:   reader.ReadFile,
}

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.Flags().BoolP("csv", "f", false, "default option reads from json input & setting this falg to true reads from CSV")
}
