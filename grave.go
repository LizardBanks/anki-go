package anki

type Grave struct {
	UpdateSequenceNumber int `db:"usn"`
	OriginalID           int `db:"oid"`
	Type                 int `db:"type"`
}
