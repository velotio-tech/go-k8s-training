package journal

import (
	"fmt"
	"os"
)

const MaxEntryLimit = 50

type Journal struct {
	Entries []entry
	File    *os.File
}

func (j *Journal) AddEntry(e *entry) error {
	if len(j.Entries) == MaxEntryLimit {
		j.Entries = append(j.Entries[1:], *e)
	} else {
		j.Entries = append(j.Entries, *e)
	}
	return UpdateJournal(j.File, j.Entries)
}

func (j *Journal) PrintList() {
	for _, entry := range j.Entries {
		fmt.Println(entry.String())
	}
}
