package cmd

import (
	"bytes"
	"fmt"
	"strings"
)

func printConsoleTent(w Tent, showYear bool) {
	if showYear {
		fmt.Printf("%v - %v\n", w.Year, w.Name)
	} else {
		fmt.Printf("%v\n", w.Name)
	}
}

func printConsoleTentInTent(w Tent) {
	for i := 0; i < len(w.Year); i++ {
		asciitent[i] = byte(w.Year[i])
	}
	firstNewLine := bytes.Index(asciitent, []byte("\n"))
	for i := 0; i < len(w.Name); i++ {
		asciitent[i+firstNewLine+1] = byte(w.Name[i])
	}
	fmt.Println(string(asciitent))
}

func printAllConsoleTents(ftw []Tent) {
	// handle if only 1 found and you want them in a tent
	if inatent {
		if len(ftw) > 1 {
			fmt.Println("too many for the pretty tent, ignoring 'inthetent' ")
		} else {
			printConsoleTentInTent(ftw[0])
			return
		}
	}

	for _, w := range ftw {
		switch sortBy {
		case "firstname", "lastname":
			names := strings.Split(w.Name, " ")

			if sortBy == "lastname" {
				fmt.Printf("%v, %v - %v\n", names[1], names[0], w.Year)
			} else {
				fmt.Printf("%v, %v - %v\n", names[0], names[1], w.Year)
			}
		case "year":
			printConsoleTent(w, true)
		}
	}
}
