package str

import "strings"

type Str struct {
	value string
}

func NewStr(s string) Str {
	return Str{value: s}
}
func StrOf(s string) *Str {
	return &Str{value: s}
}

func (s *Str) TrimSpace() *Str {
	s.value = strings.TrimSpace(s.value)
	return s
}

func (s *Str) ToString() string {
	return s.value
}

func (s *Str) Lower() *Str {
	s.value = strings.ToLower(s.value)
	return s
}

func (s *Str) Upper() *Str {
	s.value = strings.ToUpper(s.value)
	return s
}

func (s *Str) Append(sub string) *Str {
	s.value = s.value + sub
	return s
}

func (s *Str) Prepend(sub string) *Str {
	s.value = sub + s.value
	return s
}

func (s Str) Contains(substring string) bool {
	return strings.Contains(s.value, substring)
}
func (s *Str) Replace(needle string, with string) *Str {
	s.value = strings.ReplaceAll(s.value, needle, with)
	return s
}
