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
var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:     "clean",
	Aliases: []string{"c"},
	Short:   "Clean duplicates.",
	Long: `Daniel Rivas <danielrivasmd@gmail.com>

` + chalk.Green.Color("Lou") + ` will clean duplicates at a target location.
`,

	Example: `
` + chalk.Cyan.Color("lou") + ` clean
` + chalk.Cyan.Color("lou") + ` clean -l $(pwd)`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {

		// determine location
		location, ε := cmd.Flags().GetString("location")
		if ε != nil {
			log.Fatal(ε)
		}

		// execute logic
		cleanDir(location)

	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(cleanCmd)

	// flags

}

////////////////////////////////////////////////////////////////////////////////////////////////////

func matchDir(location string) {

	file, ε := os.Open(location)
	if ε != nil {
		log.Fatal(ε)
	}
	defer file.Close()

	fileList, ε := file.Readdir(0)
	if ε != nil {
		log.Fatal(ε)
	}

	// declare regex
	const ρε = `\(\d\)\w*`
	ρ := regexp.MustCompile(ρε)

	// switch
	ζ := true

	// check each file @ location
	for _, files := range fileList {
		μ := ρ.MatchString(files.Name())
		if μ {
			ζ = false
			fmt.Println(location + files.Name())
			os.Remove(location + files.Name())
		}
	}

	// trigger if no duplicates are found
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
