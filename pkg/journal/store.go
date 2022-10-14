package journal

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/velotio-tech/go-k8s-training/pkg/secure"
)

func Get(f *os.File) (*Journal, error) {
	journalBytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	entries := []entry{}
	if len(journalBytes) == 0 {
		return &Journal{
			Entries: entries,
			File:    f,
		}, nil
	}
	err = secure.Decrypt(string(journalBytes), &entries)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &Journal{
		Entries: entries,
		File:    f,
	}, nil
}

func UpdateJournal(f *os.File, journal []entry) error {
	journalBytes, err := secure.Encrypt(journal)
	if err != nil {
		return err
	}
	err = f.Truncate(0)
	if err != nil {
		return err
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(f, string(journalBytes))
	return err
}
