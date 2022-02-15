/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// mumbleCmd represents the mumble command
var mumbleCmd = &cobra.Command{
	Use:     "mumble",
	Short:   "If said softly you might find what you need",
	Aliases: []string{"search", "find", "match"},
	Long: `Find matches containing text and list them:
e.g. to find all winners in the 2010's you could just type "201":
tent mumble 201`,
	Run: func(cmd *cobra.Command, args []string) {
		winners, err := findConsoleTents(args[0], theConsoleTents, &MatchAnywhere{})
		if err != nil {
			fmt.Printf("welp: %v\n", err)
			return
		}
		printWinners(winners)
	},
}

func init() {
	rootCmd.AddCommand(mumbleCmd)

	mumbleCmd.Flags().BoolVarP(&reverse, "reverse", "r", false, "Show earlist first")
	mumbleCmd.Flags().StringVarP(&sortBy, "sort", "s", "year", "Sort by year|firstname|lastname")
}
