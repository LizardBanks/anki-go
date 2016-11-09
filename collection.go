package anki

import (
	"encoding/json"

	gorp "gopkg.in/gorp.v1"
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

func GetCollections(dbmap *gorp.DbMap) ([]*Collection, error) {
	results := []*Collection{}
	_, err := dbmap.Select(&results, "SELECT * from col")
	if err != nil {
		return make([]*Collection, 0), err
	}

	for _, collection := range results {
		// Parse JSON object
		var r map[string]*json.RawMessage
		err = json.Unmarshal([]byte(collection.DecksJSON), &r)
		if err != nil {
			return make([]*Collection, 0), err
		}

		var decks []Deck
		for _, jsonObject := range r {
			var deck Deck
			json.Unmarshal(*jsonObject, &deck)
			decks = append(decks, deck)
		}

		collection.Decks = decks
	}

	return results, nil
}
