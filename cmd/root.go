/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "local.dev"

//go:embed tent.ascii
var asciitent []byte

//go:embed consoletents.json
var winnerfile []byte

var inatent bool

var theConsoleTents Tents

type Tent struct {
	Name      string `json:"name"`
	Year      string `json:"year"`
	foundName bool
}

type Tents struct {
	Tents []Tent `json:"tents"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tent",
	Short: "Find and show the Novan console tent of the year",
	Long: `So you want to know who the console tent was for a particular year?
Perhaps you want to know which year someone was a console tent?
Or you just want to print a tent in your console?
Then look no further!
Just type the year like so:
 tent 2021
or their first name:
 tent Phil
If you want a list, well I can give you a list:
 tent list
The rest of this amazing application you can explore by yourself.
`,
	Version: version,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(defCmd string) {
	var cmdFound bool
	cmd := rootCmd.Commands()

	// check if a cmd or alias of command matches an arg
	for _, c := range cmd {
		for _, arg := range os.Args[1:] {
			if c.Name() == arg {
				cmdFound = true
				break
			}
			for _, a := range c.Aliases {
				if a == arg {
					cmdFound = true
					break
				}
			}
		}
	}
	if !cmdFound && !isHelpCommand(os.Args[1:]) && !(os.Args[1] == "completion") {
		args := append([]string{defCmd}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func isHelpCommand(cmds []string) bool {
	if len(cmds) == 0 {
		return true
	}
	return cmds[0] == "help" || cmds[0] == "-h" || cmds[0] == "--help"
}

func init() {
	json.Unmarshal(winnerfile, &theConsoleTents)
	rootCmd.PersistentFlags().BoolVarP(&inatent, "inatent", "t", false, "if you want to more tent, this is how you get more tent")
}
