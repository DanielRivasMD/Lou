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
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

// bibloCmd represents the article command
var bibloCmd = &cobra.Command{
	Use:     "biblo",
	Aliases: []string{"b"},
	Short:   "Handles all biblography operations.",
	Long: `Daniel Rivas <danielrivasmd@gmail.com>

` + chalk.Green.Color("Lou") + ` handles all biblography operations.

For example:
- Reformat articles and their references downloaded by Kopernico.
- Relocate articles manually renamed.
- Relocate theses to the proper archive.
`,

	Example: `
` + chalk.Cyan.Color("lou") + ` biblo annotate
` + chalk.Cyan.Color("lou") + ` biblo gather
` + chalk.Cyan.Color("lou") + ` biblo thesis`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	ValidArgs: []string{"format", "thesis"},
	Args:      cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// execute logic
		bibloArgs(findHome(), args)

		fmt.Println()

		// determine location
		location, _ := cmd.Flags().GetString("location")

		// clean after relocating
		cleanDir(location)

	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(bibloCmd)

	// flags

}

////////////////////////////////////////////////////////////////////////////////////////////////////

func bibloArgs(home string, args []string) {

	// allocate command
	var commd string

	// shell call
	switch args[0] {

	case "format":

		format(home, "[A-Z][a-z-]+-[0-9]{4}[A-Za-z_0-9-]+.")
		format(home, "[A-Z][a-z]+[-]{1}[A-Za-z_0-9-]+.")

	case "thesis":
		commd = home + "/Factorem/Lou/sh/thesis.sh"
		runSh(commd)

	}

}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runSh(commd string) {

	// lineBreaks
	lineBreaks()

	// buffers
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	// run command
	shCmd := exec.Command(commd)
	shCmd.Stdout = &stdout
	shCmd.Stderr = &stderr
	_ = shCmd.Run()

	// stdout
	fmt.Println(chalk.Cyan.Color(stdout.String()))

	// stderr
	if stderr.String() != "" {
		fmt.Println(chalk.Red.Color(stderr.String()))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func format(home, regString string) {

	// declare arrays
	typeArray := [2]string{"pdf", "ris"}
	folderArray := [2]string{"PDFs", "Refs"}

	// read downloads
	files, ε := ioutil.ReadDir(home + "/Downloads/")
	if ε != nil {
		log.Fatal(ε)
	}

	// loop over types
	for ix := 0; ix < len(typeArray); ix++ {

		// compile regex
		reg, _ := regexp.Compile(regString + typeArray[ix])

		// count if files are present
		τ := 0

		// loop over files
		for _, file := range files {

			// collect files
			original := reg.FindString(file.Name())

			// check for match
			if original != "" {

				// increase count
				τ++

				// trim suffix
				original = strings.TrimSuffix(original, "."+typeArray[ix])

				// define target
				target := ""
				for fx, field := range strings.Split(original, "-") {

					// accept only 7 fields
					if fx > 7 {
						break
					}

					// TODO: modify regex to cover two letter author last name
					if len(field) > 3 || fx == 0 {
						switch fx {
						case 0:
							target += field
						case 1:
							target += "-" + field
						default:
							target += "_" + field
						}
					}
				}

				// define full paths
				fullOriginal := home + "/Downloads/" + original + "." + typeArray[ix]
				fullTarget := home + "/Articulos/" + folderArray[ix] + "/" + target + "." + typeArray[ix]

				fmt.Println(chalk.Green.Color(original+"."+typeArray[ix]) + "\t\t\t" + chalk.Cyan.Color(target+"."+typeArray[ix]))

				// relocate
				ε := os.Rename(fullOriginal, fullTarget)
				if ε != nil {
					log.Fatal(ε)
				}

			}

		}

		if τ == 0 {
			emptyMessage := `
	No files to reformat:	` + folderArray[ix]

			fmt.Println(emptyMessage)
		}

	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
