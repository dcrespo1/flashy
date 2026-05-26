/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available decks",
	Run: func(cmd *cobra.Command, args []string) {
		err := ListDecks()
		if err != nil {
			fmt.Println("error listing decks: ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
