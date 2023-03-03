package cmd

import (
	"os"

	"github.com/TheMorgz/gobirth/internal/output"
	"github.com/spf13/cobra"
)

var printer output.Printer

var rootCmd = &cobra.Command{
	Use:   "gobirth",
	Short: "A program to print out the people whose birthday is today.",
	Long:  `Given a JSON file with a list of people and their dates of birth, write a program to print out the people whose birthday is today.`,
}

func Execute() {
	printer.ClearScreen()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
