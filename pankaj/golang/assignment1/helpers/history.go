package helpers

import "container/list"

type History struct {
	Command *list.List
	size    int
}

func (h *History) Init(size int) {
	h.Command = list.New()
	h.size = size
}

func (h *History) Capture(cmd string) {
	h.Command.PushFront(cmd)
	if h.Command.Len() > h.size {
		h.Command.Remove(h.Command.Back())
	}
}
