package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func SaveDeck(deck Deck) error{
	fileData, err := json.MarshalIndent(deck, "", " ")
	if err != nil {
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
    	return err
	}
	targetDir := filepath.Join(homeDir, ".flashy", "decks")

	err = os.MkdirAll(targetDir, 0755)
	if err != nil {
		return err
	}

	fullPath := filepath.Join(targetDir, deck.DeckName + ".json")

	err = os.WriteFile(fullPath, fileData, 0644)
	if err != nil{
		return err
	}

	return nil
}

func LoadDeck(name string) (Deck, error){
	homeDir, err := os.UserHomeDir()
	if err != nil {
    	return Deck{}, err
	}
	targetDir := filepath.Join(homeDir, ".flashy", "decks")
	fullPath := filepath.Join(targetDir, name + ".json")

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