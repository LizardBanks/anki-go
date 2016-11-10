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
ankiClient, err := anki.NewClient(PATH)
if err != nil {
	log.Fatalf("%v", err)
}

collections, err := ankiClient.GetCollections()
if err != nil {
	log.Fatalf("Error: %v\n", err)
}

log.Printf("Number of collections: %v\n", len(collections))

for _, collection := range collections {
	for _, deck := range collection.Decks {
		log.Printf("Deck Name: %v\n", deck.Name)
	}
}
```