package cmd

type Card struct {
	Question string `json:"question"`
	Answer string `json:"answer"`
	Correct *bool `json:"correct,omitempty"`
}

type Deck struct {
    DeckName string `json:"deck_name"`
    Cards    []Card `json:"cards"`
}

type StudySession struct {
	StudiedDecks Deck
	NumberCorrect int
}


// func CreateDeck(name string) Deck {  // this works but can be simplified
// 	cards := []Card{}
// 	deck := Deck{
// 		DeckName: name,
// 		Cards: cards,
// 	}

// 	return deck
// }

func CreateDeck(name string) Deck{
	return Deck{
		DeckName: name,
		Cards: []Card{},
	}
}