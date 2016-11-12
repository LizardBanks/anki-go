package anki

import "fmt"

const (
	CardTypeNew          = 0
	CardTypeLearning     = 1
	CardTypeDue          = 2
	QueueTypeSuspended   = -1
	QueueTypeUserBuried  = -2
	QueueTypeSchedBuried = -3
)

type Card struct {
	ID                   int    `db:"id"`
	NoteID               int    `db:"nid"`
	DeckID               int    `db:"did"`
	Ordinal              int    `db:"ord"`
	ModificationTime     int    `db:"mod"`
	UpdateSequenceNumber int    `db:"usn"`
	Type                 int    `db:"type"`
	Queue                int    `db:"queue"`
	Due                  int    `db:"due"`
	Interval             int    `db:"ivl"`
	Factor               int    `db:"factor"`
	Reps                 int    `db:"reps"`
	Lapses               int    `db:"lapses"`
	Left                 int    `db:"left"`
	OriginalDue          int    `db:"odue"`
	OriginalDid          int    `db:"odid"`
	Flags                int    `db:"flags"`
	Data                 string `db:"data"`
}

func (client *Client) GetCards() ([]Card, error) {
	results := []Card{}
	_, err := client.DBHandle.Select(&results, "SELECT * from cards")
	if err != nil {
		return make([]Card, 0), err
	}

	return results, nil
}

func (client *Client) GetCardsForDeck(deck *Deck) ([]Card, error) {
	results := []Card{}
	query := fmt.Sprintf("SELECT * from cards WHERE did = %d", deck.ID)
	_, err := client.DBHandle.Select(&results, query)
	if err != nil {
		return make([]Card, 0), err
	}

	return results, nil
}
