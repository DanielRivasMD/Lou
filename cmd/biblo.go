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
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	Short:   "Handle all biblography operations.",
	Long: `Daniel Rivas <danielrivasmd@gmail.com>

` + chalk.Green.Color("Lou") + ` handles all biblography operations.

For example:
- Annotate articles.
- Build library from ` + chalk.Cyan.Color("nbib") + `.
- Reformat articles and their references downloaded by ` + chalk.Yellow.Color("Kopernico") + `.
- Relocate articles manually renamed.
- Relocate theses to the proper archive.
`,

	Example: `
` + chalk.Cyan.Color("lou") + ` biblo annotate
` + chalk.Cyan.Color("lou") + ` biblo build
` + chalk.Cyan.Color("lou") + ` biblo gather
` + chalk.Cyan.Color("lou") + ` biblo thesis`,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func bibloArgs(home string) {

	// search for downloads from kopernico
	format(home, "[A-Z][a-z-]+-[0-9]{4}[A-Za-z_0-9-]+.")

	// search for manual renamed
	format(home, "[A-Z][a-z]+[-]{1}[A-Za-z_0-9-]+.")

}

////////////////////////////////////////////////////////////////////////////////////////////////////

func format(home, ρε string) {

	// declare arrays
	typeArray := [2]string{"pdf", "ris"}
	folderArray := [2]string{"PDFs", "Refs"}

	// read downloads
	files, ε := ioutil.ReadDir(home + "/Downloads/")
	if ε != nil {
		log.Fatal(ε)
	}

	// loop over types
	for ι := 0; ι < len(typeArray); ι++ {

		// compile regex
		ρ, ε := regexp.Compile(ρε + typeArray[ι])
		if ε != nil {
			log.Fatal(ε)
		}

		// count if files are present
		τ := 0

		// loop over files
		for _, ƒ := range files {

			// collect files
			original := ρ.FindString(ƒ.Name())

			// check for match
			if original != "" {

				// increase count
				τ++

				// trim suffix
				original = strings.TrimSuffix(original, "."+typeArray[ι])

				// define target
				target := ""
				for ο, field := range strings.Split(original, "-") {

					// accept only 7 fields
					if ο > 7 {
						break
					}

					if len(field) > 3 || ο == 0 {
						switch ο {
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
				fullOriginal := home + "/Downloads/" + original + "." + typeArray[ι]
				fullTarget := home + "/Articulos/" + folderArray[ι] + "/" + target + "." + typeArray[ι]

				fmt.Println(chalk.Green.Color(original+"."+typeArray[ι]) + "\t\t\t" + chalk.Cyan.Color(target+"."+typeArray[ι]))

				// relocate
				ε := os.Rename(fullOriginal, fullTarget)
				if ε != nil {
					log.Fatal(ε)
				}

			}

		}

		if τ == 0 {
			emptyMessage := `
	No files to reformat:	` + folderArray[ι]

			fmt.Println(emptyMessage)
		}

	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(bibloCmd)

	// flags

}

////////////////////////////////////////////////////////////////////////////////////////////////////
