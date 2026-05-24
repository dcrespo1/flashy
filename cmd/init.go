/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize flashy configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if ConfigExists() {
			fmt.Println("Config already exists, overwrite? (y/n)")
			var response string
			fmt.Scan(&response)
			if response == "n" {
				fmt.Println("Keeping existing config.")
            	return
			}
		}
		fmt.Print("Enter decks directory (default: ~/.flashy/decks): ")
    	scanner := bufio.NewScanner(os.Stdin)
    	scanner.Scan()
		input := scanner.Text()

		if input == "" {
			defaultDir, err := defaultDecksDir()
			if err != nil {
				fmt.Println("Error getting default directory:", err)
        		return
			}
			input = defaultDir
		}

    	cfg := Config{
        	DecksDir: input,
    	}
		err := SaveConfig(cfg)
		if err != nil {
			fmt.Println("Error Saving Config:", err)
			return
		}
		fmt.Println("Config Saved!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
