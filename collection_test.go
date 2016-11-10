package anki

import (
	"fmt"
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

	if r[0].Name != "Default" {
		t.Errorf("expected first deck to be 'Default', got '%s'", r[0].Name)
	}

	if r[1].Name != "Korean" {
		t.Errorf("expected first deck to be 'Korean', got '%s'", r[1].Name)
	}

	testJSONString = "\"1\": {\"foo\": \"Default\"}, \"1234\": {\"name\": \"Korean\"}}"
	r, err = parseDecksJSON(testJSONString)

	if err == nil {
		fmt.Printf("expected error for invalid JSON")
	}
}
