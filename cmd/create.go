package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [deck name]",
	Short: "Create a new flashcard deck",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		joinedArgs := strings.Join(args, " ")
		deck := CreateDeck(joinedArgs)
		err := SaveDeck(deck)
		if err != nil {
    		fmt.Println("Error saving deck:", err)
    		return
		}
		fmt.Println("Created Deck:", deck.DeckName)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
