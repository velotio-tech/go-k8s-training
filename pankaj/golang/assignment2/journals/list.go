package journals

import "container/list"

type Journal struct {
	Item *list.List
	size int
}

func (journal *Journal) Init(size int) {
	journal.Item = list.New()
	journal.size = size
}

func (journal *Journal) Capture(cmd string) {
	journal.Item.PushBack(cmd)
	if journal.Item.Len() > journal.size {
		journal.Item.Remove(journal.Item.Front())
	}
}
