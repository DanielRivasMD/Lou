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
	"os"

	"github.com/atrox/homedir"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

var rootCmd = &cobra.Command{
	Use:     "lou",
	Version: "v0.2",
	Short:   chalk.Green.Color("Lou") + ", personal assitant at your service",
	Long: `Daniel Rivas <danielrivasmd@gmail.com>

` + (chalk.Green.Color("Lou") + chalk.White.Color(", personal assitant at your service")),

	Example: `
` + chalk.Cyan.Color("lou") + ` help`,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func Execute() {
	ε := rootCmd.Execute()
	if ε != nil {
		fmt.Println(ε)
		os.Exit(1)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	cobra.OnInitialize(initConfig)

	// persistent flags
	rootCmd.PersistentFlags().StringP("location", "l", "/Users/drivas/Downloads/", "Location to clean")

}

////////////////////////////////////////////////////////////////////////////////////////////////////

func initConfig() {

}

////////////////////////////////////////////////////////////////////////////////////////////////////

// find home directory
func findHome() string {
	home, ε := homedir.Dir()
	if ε != nil {
		fmt.Println(ε)
		os.Exit(1)
	}
	return home
}

////////////////////////////////////////////////////////////////////////////////////////////////////
