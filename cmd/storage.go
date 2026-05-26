package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func SaveDeck(deck Deck) error{
	fileData, err := json.MarshalIndent(deck, "", " ")
	if err != nil {
		return err
	}

	cfg, err := LoadConfig()
	if err != nil {
		return err
	}
	err = os.MkdirAll(cfg.DecksDir, 0755)
	if err != nil{
		return err
	}

	fullPath := filepath.Join(cfg.DecksDir, deck.DeckName + ".json")

	err = os.WriteFile(fullPath, fileData, 0644)
	if err != nil{
		return err
	}

	return nil
}

func LoadDeck(name string) (Deck, error){
	cfg, err := LoadConfig()
	if err != nil {
		return Deck{}, err
	}

	fullPath := filepath.Join(cfg.DecksDir, name + ".json")

	data, err := os.ReadFile(fullPath)
	if err != nil {
		return Deck{}, err
	}

	var deck Deck
	if err = json.Unmarshal(data, &deck); err != nil {
		return Deck{}, err
	}
	return deck, err

}

func ListDecks() error {
	cfg, err := LoadConfig()
	if err != nil {
		return err
	}
	
	deckList, err := os.ReadDir(cfg.DecksDir)
	if err != nil {
		return err
	}

	if len(deckList) == 0 {
    	fmt.Println("No decks found. Create one with: flashy create <name>")
    	return nil
	}else {
		for _, entry := range deckList {
			name := strings.TrimSuffix(entry.Name(), ".json")
			fmt.Println(name)

		}
		return nil
	}

}