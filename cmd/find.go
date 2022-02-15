package cmd

import (
	"fmt"
	"strings"
)

func getConsoleTent(search string) (*Tent, error) {
	cts, _ := findConsoleTents(search, theConsoleTents, nil)
	if len(cts) > 1 {
		return nil, fmt.Errorf("a few too many tents and no exact match, try to mumble and i'm sure they'll all respond")
	}

	if len(cts) == 1 {
		return &cts[0], nil
	}

	return nil, fmt.Errorf("hmmm can't find %q, perhaps try mumbling and see who responds?", search)
}

func findConsoleTents(search string, cts Tents, match matchInterface) ([]Tent, error) {
	// set default match
	if match == nil {
		match = &matchExact{}
	}

	var ftw []Tent
	for _, v := range cts.Tents {
		names := strings.Split(v.Name, " ")
		for _, n := range names {
			if match.match(n, search) {
				v.foundName = true
				ftw = append(ftw, v)
			}
		}
		if match.match(v.Year, search) {
			ftw = append(ftw, v)
		}
	}
	if len(ftw) > 0 {
		return ftw, nil
	}
	return nil, fmt.Errorf("hmmm can't find %q, they must have left the camp", search)
}
