package cmd

import (
	"sort"
	"strings"
)

var reverse bool
var sortBy string

type byYear []Tent

func (w byYear) Len() int {
	return len(w)
}

func (w byYear) Less(i, j int) bool {
	if reverse {
		return w[i].Year > w[j].Year
	}
	return w[i].Year < w[j].Year
}

func (w byYear) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

type byName []Tent

func (w byName) Len() int {
	return len(w)
}

func (w byName) Less(i, j int) bool {
	iname := strings.Split(w[i].Name, " ")
	jname := strings.Split(w[j].Name, " ")
	iindex := 0
	jindex := 0
	if sortBy == "lastname" {
		iindex = len(iname) - 1
		jindex = len(jname) - 1
	}
	if reverse {
		return iname[iindex] > jname[jindex]
	}
	return iname[iindex] < jname[jindex]
}

func (w byName) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func sortWinners(w Tents) Tents {
	switch sortBy {
	case "firstname", "lastname":
		sort.Sort(byName(w.Tents))
	case "year":
		sort.Sort(byYear(w.Tents))
	}

	return w
}
