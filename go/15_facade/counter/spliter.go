package counter

import (
	"strings"
)

type spliter struct {
	sep string
}

func (s *spliter) split(sentence string) []string {
	return strings.Split(sentence, s.sep)
}