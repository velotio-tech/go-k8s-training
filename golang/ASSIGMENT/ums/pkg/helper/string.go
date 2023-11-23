package helper

import (
	"strings"
	"unicode"
)

/*
Author: Arijit Nayak
Description: RemoveAllUnNecessarySpaces removes all un-necessary spaces.
             This method is added due to the best time and space management.
             BenchMark Data -> Operation Time - 932298 ns/op | Space - 1 allocs/op
*/
// RemoveAllUnNecessarySpaces ...
func (h *UserServiceHelper) RemoveAllUnNecessarySpaces(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
