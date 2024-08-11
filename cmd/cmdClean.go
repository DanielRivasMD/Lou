/*
Copyright © 2020 Daniel Rivas <danielrivasmd@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
const Я = `\(\d\)\w*` // backticks are used here to contain the expression

// declare regex
var (
	ρ = regexp.MustCompile(Я)
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// cleanCmd
var cleanCmd = &cobra.Command{
	Use:     "clean",
	Aliases: []string{"c"},
	Short:   "Clean duplicates.",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` + chalk.Green.Color("Lou") + ` will clean duplicates at a target location.
`,

	Example: `
` + chalk.Cyan.Color("lou") + ` clean
` + chalk.Cyan.Color("lou") + ` clean -l $(pwd)`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {
		// determine location
		location, _ := κ.Flags().GetString("location")

		// execute logic
		cleanDir(location)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO: migrate functions to files
func matchDir(location string) {
	directory, ε := os.Open(location)
	if ε != nil {
		log.Fatal(ε)
	}
	defer directory.Close()

	ł, ε := directory.Readdir(0)
	if ε != nil {
		log.Fatal(ε)
	}

	// switch
	ζ := true

	// check each file @ location
	for _, ƒ := range ł {
		if ρ.MatchString(ƒ.Name()) {
			ζ = false
			fmt.Println(location + ƒ.Name())
			os.Remove(location + ƒ.Name())
		}
	}

	// trigger if no duplicates found
	if ζ {
		fmt.Println(chalk.Cyan.Color("\tNo files to remove"))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func cleanDir(location string) {
	// lineBreaks
	lineBreaks()

	// function call
	matchDir(location)

	// buffers
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	// stdout
	fmt.Println(chalk.Cyan.Color(stdout.String()))

	// stderr
	if stderr.String() != "" {
		fmt.Println(chalk.Red.Color(stderr.String()))
	}

	// lineBreaks
	lineBreaks()
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(cleanCmd)

	// flags
}

////////////////////////////////////////////////////////////////////////////////////////////////////
