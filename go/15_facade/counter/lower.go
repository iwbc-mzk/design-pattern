package counter

import (
	"strings"
)

type lower struct {}

func (l *lower) lower(sentence string) string {
	return strings.ToLower(sentence)
}