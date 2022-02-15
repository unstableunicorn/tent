/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// yellCmd represents the yell command
var yellCmd = &cobra.Command{
	Use:     "yell",
	Aliases: []string{"list", "ls"},
	Short:   "Show me all the console tents",
	Long: `If you just want a list with various sorting abilities this is it.
tent yell
tent yell --reverse
tent yell --sort lastname`,
	Run: func(cmd *cobra.Command, args []string) {
		winners := sortWinners(theConsoleTents)
		printWinners(winners.Tents)
	},
}

func init() {
	rootCmd.AddCommand(yellCmd)

	yellCmd.Flags().BoolVarP(&reverse, "reverse", "r", false, "Show earlist first")
	yellCmd.Flags().StringVarP(&sortBy, "sort", "s", "year", "Sort by year|firstname|lastname")
}
