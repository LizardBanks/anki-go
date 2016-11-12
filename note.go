package anki

type Note struct {
	ID                   int    `db:"id"`
	GlobalUniqueID       string `db:"guid"`
	ModelID              int    `db:"mid"`
	ModificationTime     int    `db:"mod"`
	UpdateSequenceNumber int    `db:"usn"`
	Tags                 string `db:"tags"`
	Fields               string `db:"flds"`
	SortField            string `db:"sfld"`
	Checksum             int    `db:"csum"`
	Flags                int    `db:"flags"`
	Data                 string `db:"data"`
}
