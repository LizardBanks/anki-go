package anki

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

// Create mock struct for earier mocking of select functions
type mockHandle struct {
	SelectFun func(i interface{}, query string, args ...interface{}) ([]interface{}, error)
}

func (handle *mockHandle) Select(i interface{}, query string, args ...interface{}) ([]interface{}, error) {
	return handle.SelectFun(i, query, args)
}

func Test_parseDecksJSON(t *testing.T) {
	testJSONString := "{\"1\": {\"name\": \"Default\"}, \"1234\": {\"name\": \"Korean\"}}"
	r, err := parseDecksJSON(testJSONString)

	if err != nil {
		t.Error("expected no error")
	}

	if len(r) != 2 {
		t.Errorf("expected 2 decks, got %d", len(r))
	}

	expectedStrings := []string{"Default", "Korean"}
	for _, expectedString := range expectedStrings {
		found := false
		for _, result := range r {
			if expectedString == result.Name {
				found = true
				break
			}
		}

		if found != true {
			t.Fatalf("expected to find '%s' but couldnt find it", expectedString)
		}
	}

	testJSONString = "\"1\": {\"foo\": \"Default\"}, \"1234\": {\"name\": \"Korean\"}}"
	r, err = parseDecksJSON(testJSONString)

	if err == nil {
		fmt.Printf("expected error for invalid JSON")
	}
}

func Test_GetCollections_DBError(t *testing.T) {
	m := mockHandle{}
	m.SelectFun = func(i interface{}, query string, args ...interface{}) ([]interface{}, error) {
		return []interface{}{}, errors.New("database exploded")
	}

	client := Client{}
	client.DBHandle = &m

	result, err := client.GetCollections()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if len(result) != 0 {
		t.Fatalf("did not expect a result, got %d", len(result))
	}
}

func Test_GetCollections_NoDeckJSON(t *testing.T) {
	m := mockHandle{}
	m.SelectFun = func(i interface{}, query string, args ...interface{}) ([]interface{}, error) {
		collection := Collection{}
		collection.ID = 1337

		fakeCollections := []*Collection{&collection}

		source := reflect.ValueOf(fakeCollections)
		dest := reflect.ValueOf(i).Elem()

		// write to destination
		dest.Set(source)

		return []interface{}{}, nil
	}

	client := Client{}
	client.DBPath = "foo"
	client.DBHandle = &m

	col, err := client.GetCollections()
	if err != nil {
		t.Fatalf("did not expect error, got %v", err)
	}

	if len(col) != 1 {
		t.Fatalf("expected 1 collection, got %d", len(col))
	}

	if len(col[0].Decks) != 0 {
		t.Fatalf("expected 0 decks in collection, got %d", len(col[0].Decks))
	}
}

func Test_GetCollections_JSONParsing(t *testing.T) {
	m := mockHandle{}
	m.SelectFun = func(i interface{}, query string, args ...interface{}) ([]interface{}, error) {
		collection := Collection{}
		collection.ID = 1337
		collection.DecksJSON = "{\"1\": {\"name\": \"Default\"}, \"1234\": {\"name\": \"Korean\"}}"

		fakeCollections := []*Collection{&collection}

		source := reflect.ValueOf(fakeCollections)
		dest := reflect.ValueOf(i).Elem()

		// write to destination
		dest.Set(source)

		return []interface{}{}, nil
	}

	client := Client{}
	client.DBPath = "foo"
	client.DBHandle = &m

	col, err := client.GetCollections()
	if err != nil {
		t.Fatalf("did not expect error, got %v", err)
	}

	if len(col) != 1 {
		t.Fatalf("expected 1 collection, got %d", len(col))
	}

	if len(col[0].Decks) != 2 {
		t.Fatalf("expected 2 decks in collection, got %d", len(col[0].Decks))
	}
}

func Test_GetDecks(t *testing.T) {
	m := mockHandle{}
	m.SelectFun = func(i interface{}, query string, args ...interface{}) ([]interface{}, error) {
		collection1 := Collection{}
		collection1.ID = 1337
		collection1.DecksJSON = "{\"1\": {\"name\": \"Default\"}, \"1234\": {\"name\": \"Korean\"}}"

		collection2 := Collection{}
		collection2.ID = 1339
		collection2.DecksJSON = "{\"888\": {\"name\": \"German\"}, \"1234\": {\"name\": \"More Korean\"}}"

		fakeCollections := []*Collection{&collection1, &collection2}

		source := reflect.ValueOf(fakeCollections)
		dest := reflect.ValueOf(i).Elem()

		// write to destination
		dest.Set(source)

		return []interface{}{}, nil
	}

	client := Client{}
	client.DBPath = "foo"
	client.DBHandle = &m

	decks, err := client.GetDecks()
	if err != nil {
		t.Errorf("did not expect error, got: '%s'", err.Error())
	}

	if len(decks) != 4 {
		t.Errorf("expected 4 decks, found %d", len(decks))
	}

	expectedStrings := []string{"Default", "German", "Korean", "More Korean"}
	for _, expectedString := range expectedStrings {
		found := false
		for _, result := range decks {
			if expectedString == result.Name {
				found = true
				break
			}
		}

		if found != true {
			t.Fatalf("expected to find '%s' but couldnt find it", expectedString)
		}
	}

}
