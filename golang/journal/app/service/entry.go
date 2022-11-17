package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/velotio-ajaykumbhar/journal/app/model"
	"github.com/velotio-ajaykumbhar/journal/util"
)

var (
	JournalFilename = "journal.lock"
	maxEntryAllowed = 50
)

func CreateEntry(title string) {

	if title == "" {
		fmt.Println("please provide input")
		return
	}

	user := GetSession()
	if user.UserId == "" {
		fmt.Println("please login or create new account.")
		return
	}

	filename := user.UserId + ".lock"
	util.InitFile(filename)

	data, err := util.ReadFile(filename)
	if err != nil {
		fmt.Println("error while add entry. please try again", err)
		return
	}

	data = util.Decode(string(data))
	entries := []model.Entry{}
	json.Unmarshal(data, &entries)

	if len(entries) >= maxEntryAllowed {
		entries = entries[1:]
	}

	entry := model.Entry{
		UserId:    user.UserId,
		Title:     title,
		Timestamp: time.Now().Format("2006-01-02 3:4:5 pm"),
	}

	entries = append(entries, entry)

	marshal, _ := json.Marshal(entries)

	encode := util.Encode(marshal)

	err = util.WriteFile(filename, []byte(encode))
	if err != nil {
		fmt.Println("error while sign up. please try again")
	}
}

func GetAllEntry() {

	user := GetSession()
	if user.UserId == "" {
		fmt.Println("please login or create new account.")
		return
	}

	filename := user.UserId + ".lock"
	util.InitFile(filename)

	data, err := util.ReadFile(filename)
	if err != nil {
		fmt.Println("error while add entry. please try again")
		return
	}

	data = util.Decode(string(data))
	entries := []model.Entry{}
	json.Unmarshal(data, &entries)

	for _, entry := range entries {
		fmt.Printf("[%s] %s\n", entry.Timestamp, entry.Title)
	}

}
