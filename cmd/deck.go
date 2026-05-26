type Rating int

const (
    Unrated   Rating = 0
    Know      Rating = 1
    NeedsWork Rating = 2
)

type Card struct {
    Question string `json:"question"`
    Answer   string `json:"answer"`
}

type Deck struct {
    DeckName string `json:"deck_name"`
    Cards    []Card `json:"cards"`
}

type StudySession struct {
    Deck           Deck
    Ratings        map[int]Rating
    TotalCards     int
    KnowCount      int
    NeedsWorkCount int
}