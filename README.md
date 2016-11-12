# anki-go

[![Build Status](https://travis-ci.org/dvcrn/anki-go.svg?branch=master)](https://travis-ci.org/dvcrn/anki-go)
[![Coverage Status](https://coveralls.io/repos/github/dvcrn/anki-go/badge.svg?branch=master)](https://coveralls.io/github/dvcrn/anki-go?branch=master)

**Very early** attempt to create a go client to interact with the [Anki](http://ankisrs.net) DB in Go. 

## What is working?

- Reading Collections
- Reading Decks
- Reading Cards

and not really anything else yet 😇

## Usage

```go
const PATH = "path/to/anki/db"

ankiClient, err := anki.NewClient(PATH)
if err != nil {
	log.Fatalf("%v", err)
}

deck, err := ankiClient.GetDeckByName("My::Deck")
if err != nil {
    log.Fatalf("err: %v\n", err)
}

cards, err := ankiClient.GetCardsForDeck(&deck)
if err != nil {
    log.Fatalf("err: %v\n", err)
}

for _, card := range cards {
    fmt.Printf("card: %v\n", card)
}
```
