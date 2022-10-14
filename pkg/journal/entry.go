package journal

import (
	"fmt"
	"strings"
	"time"
)

type entry struct {
	Timestamp time.Time `json:"timestamp"`
	Data      string    `json:"data"`
}

func NewEntry(data string) *entry {
	return &entry{
		Data:      data,
		Timestamp: time.Now(),
	}
}

func (e *entry) Validate() error {
	if strings.TrimSpace(e.Data) == "" {
		return fmt.Errorf("invalid journal entry")
	}
	return nil
}

func (e *entry) String() string {
	hour := e.Timestamp.Local().Hour() % 12
	period := ""
	if e.Timestamp.Local().Hour()/12 == 1 {
		period = "pm"
	} else {
		period = "am"
	}
	return fmt.Sprintf("%d %s %d %02d.%02d %s - %s", e.Timestamp.Local().Day(), e.Timestamp.Local().Month().String()[:3], e.Timestamp.Local().Year(), hour, e.Timestamp.Local().Minute(), period, e.Data)
}
