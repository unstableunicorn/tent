package cmd

import "strings"

type matchInterface interface {
	match(string, string) bool
}

type matchExact struct{}

func (me *matchExact) match(s1 string, s2 string) bool {
	return strings.EqualFold(s1, s2)
}

type matchAnywhere struct{}

func (ma *matchAnywhere) match(s string, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
