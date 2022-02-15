package cmd

import "strings"

type MatchInterface interface {
	Match(string, string) bool
}

type MatchExact struct{}

func (me *MatchExact) Match(s1 string, s2 string) bool {
	return strings.EqualFold(s1, s2)
}

type MatchAnywhere struct{}

func (ma *MatchAnywhere) Match(s string, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
