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

	"github.com/DanielRivasMD/Lou/aux"
	"github.com/ttacon/chalk"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:     "clean",
	Aliases: []string{"c"},
	Short:   "Lou cleans duplicates",
	Long: `Daniel Rivas <danielrivasmd@gmail.com>

Lou cleans duplicates at a target location`,
	Example: `
Lou clean
Lou clean -l $(pwd)`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {

		location, _ := cmd.Flags().GetString("location")

		// lineBreaks
		aux.LineBreaks()

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
		aux.LineBreaks()

	},

	////////////////////////////////////////////////////////////////////////////////////////////////////

}

func init() {
	rootCmd.AddCommand(cleanCmd)

	////////////////////////////////////////////////////////////////////////////////////////////////////

	// flags
	cleanCmd.Flags().StringP("location", "l", "/Users/drivas/Downloads/", "Location to clean")

	////////////////////////////////////////////////////////////////////////////////////////////////////

}

////////////////////////////////////////////////////////////////////////////////////////////////////

func matchDir(location string) {

	////////////////////////////////////////////////////////////////////////////////////////////////////

	file, err := os.Open(location)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err) // TODO: get a better error handler with package error
	}
	defer file.Close()

	fileList, _ := file.Readdir(0)

	// declare regex
	reg := regexp.MustCompile(`\(\d\)\w*`)

	// switch
	sw := true

	// check each file @ location
	for _, files := range fileList {
		m := reg.MatchString(files.Name())
		if m {
			sw = false
			fmt.Println(location + files.Name())
			os.Remove(location + files.Name())
		}
	}

	// trigger if no duplicates are found
	if sw {
		fmt.Println(chalk.Cyan.Color("\tNo files to remove"))
	}

}

////////////////////////////////////////////////////////////////////////////////////////////////////
