package helpers

import "container/list"

type History struct {
	Commands *list.List
	size     int
}

func (hist *History) Init(size int) {
	// Initializes a history object

	hist.Commands = list.New()
	hist.size = size
}

func (hist *History) Record(cmd string) {
	// Records a command in the history

	hist.Commands.PushFront(cmd)

	if hist.Commands.Len() > hist.size {
		hist.Commands.Remove(hist.Commands.Back())
	}
}
