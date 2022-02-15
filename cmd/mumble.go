/*Package cmd functions for talking softly and seeing who reponds
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
		winners, err := findConsoleTents(args[0], theConsoleTents, &matchAnywhere{})
		if err != nil {
			fmt.Printf("welp: %v\n", err)
			return
		}
		printAllConsoleTents(winners)
	},
}

func init() {
	rootCmd.AddCommand(mumbleCmd)

	mumbleCmd.Flags().BoolVarP(&reverse, "reverse", "r", false, "Show earlist first")
	mumbleCmd.Flags().StringVarP(&sortBy, "sort", "s", "year", "Sort by year|firstname|lastname")
}
