package model

type Entry struct {
	Title     string `json:"title"`
	Timestamp string `json:"timestamp"`
	UserId    string `json:"userId"`
}
