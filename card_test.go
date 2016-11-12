package anki

import (
	"strings"
	"testing"
)

func Test_GetCardsForDeck_filterByDeckID(t *testing.T) {
	m := mockHandle{}
	m.SelectFun = func(i interface{}, query string, args ...interface{}) ([]interface{}, error) {
		if !strings.Contains(query, "did = 1337") {
			t.Errorf("query does not contain 'did = ' filter")
		}

		return []interface{}{}, nil
	}

	client := Client{}
	client.DBHandle = &m

	deck := Deck{
		ID: 1337,
	}

	client.GetCardsForDeck(&deck)
}
