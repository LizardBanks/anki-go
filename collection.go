package anki

import (
	"encoding/json"
	"errors"
	"fmt"
)

const collectionTableName = "col"

type Deck struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	ExtendedRev          int    `json:"extendedRev"`
	UpdateSequenceNumber int    `json:"usn"`
	Collapsed            bool   `json:"collapsed"`
	BrowserCollapsed     bool   `json:"browserCollapsed"`
	NewToday             []int  `json:"newToday"`
	TimeToday            []int  `json:"timeToday"`
	Dynamic              int    `json:"dyn"`
	ExtendedNew          int    `json:"extendedNew"`
	Conf                 int    `json:"conf"`
	ReviewToday          []int  `json:"revToday"`
	LearnToday           []int  `json:"lrnToday"`
	Mod                  int    `json:"mod"`
	Desc                 string `json:"desc"`
}

type Collection struct {
	ID                   int `db:"id"`
	Created              int `db:"crt"`
	LastModified         int `db:"mod"`
	SchemaModTime        int `db:"scm"`
	Version              int `db:"ver"`
	Dirty                int `db:"dty"`
	UpdateSequenceNumber int `db:"usn"`
	LastSyncTime         int `db:"ls"`

	ConfigurationJSON string `db:"conf"`
	ModelsJSON        string `db:"models"`
	DecksJSON         string `db:"decks"`
	DeckOptionsJSON   string `db:"dconf"`

	Decks []Deck `db:"-"`

	Tags string `db:"tags"`
}

func parseDecksJSON(jsonString string) ([]Deck, error) {
	var r map[string]*json.RawMessage
	err := json.Unmarshal([]byte(jsonString), &r)
	if err != nil {
		return make([]Deck, 0), err
	}

	var decks []Deck
	for _, jsonObject := range r {
		var deck Deck
		json.Unmarshal(*jsonObject, &deck)
		decks = append(decks, deck)
	}

	return decks, nil
}

func (client *Client) GetCollections() ([]*Collection, error) {
	results := []*Collection{}
	_, err := client.DBHandle.Select(&results, "SELECT * from col")
	if err != nil {
		return make([]*Collection, 0), err
	}

	for _, collection := range results {
		decks, err := parseDecksJSON(collection.DecksJSON)

		// on error, just assign empty decklist
		// since we always get one back, just ignore the error
		if err != nil {
			fmt.Printf("error in deckset: %s\n", err.Error())
		}

		collection.Decks = decks
	}

	return results, nil
}

func (client *Client) GetDecks() ([]Deck, error) {
	collections, err := client.GetCollections()
	if err != nil {
		return []Deck{}, err
	}

	var decks []Deck
	for _, collection := range collections {
		for _, deck := range collection.Decks {
			decks = append(decks, deck)
		}
	}

	return decks, nil
}

func (client *Client) GetDeckByName(name string) (Deck, error) {
	decks, err := client.GetDecks()
	if err != nil {
		return Deck{}, err
	}

	for _, deck := range decks {
		if deck.Name == name {
			return deck, nil
		}
	}

	return Deck{}, errors.New("couldn't find deck with name")
}
