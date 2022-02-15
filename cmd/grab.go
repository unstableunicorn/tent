/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// grabCmd represents the grab command
var grabCmd = &cobra.Command{
	Use:     "grab",
	Aliases: []string{"get", "want", "question"},
	Short:   "Grab the person",
	Long:    `Like find, but you know what you are after`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		w, err := getConsoleTent(args[0])
		if err != nil {
			fmt.Printf("tent is empty: %v\n", err)
			return
		}
		if inatent {
			printWinnerInTent(*w)
		} else {
			printWinner(*w, w.foundName)
		}
	},
}

func init() {
	rootCmd.AddCommand(grabCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grabCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grabCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
