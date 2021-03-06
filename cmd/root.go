/*Package cmd root function for console tents
Copyright © 2022 Elric Hindy <anunstableunicorn@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	// blank as only needed for FS embedding
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

// Tent stores the Console Tents Details
type Tent struct {
	Name      string `json:"name"`
	Year      string `json:"year"`
	foundName bool
}

// Tents holds more than one Tent
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
	if !cmdFound && !isRootCommand(os.Args[1:]) {
		args := append([]string{defCmd}, os.Args[1:]...)
		rootCmd.SetArgs(args)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func isRootCommand(cmds []string) bool {
	excludedCommands := []string{
		"help",
		"-h",
		"--help",
		"-v",
		"--version",
		"completion",
	}

	if len(cmds) == 0 {
		return true
	}

	for _, v := range excludedCommands {
		if v == cmds[0] {
			return true
		}
	}
	return false
}

func init() {
	json.Unmarshal(winnerfile, &theConsoleTents)
	rootCmd.PersistentFlags().BoolVarP(&inatent, "inatent", "t", false, "if you want to more tent, this is how you get more tent")
}
