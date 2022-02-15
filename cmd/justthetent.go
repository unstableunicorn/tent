/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// justthetentCmd represents the iwanttent command
var justthetentCmd = &cobra.Command{
	Use:     "justthetent",
	Aliases: []string{"iwanttent", "tentytenttent", "justtent", "nothingbuttent", "showmethetent", "ohlookatent"},
	Short:   "Just show me the tent",
	Long: `Sometimes we just want to see the tent.
Sometimes we desire the tent.
This will fill your desire.
This will print a tent in your console`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(string(asciitent))
	},
}

func init() {
	rootCmd.AddCommand(justthetentCmd)
}
