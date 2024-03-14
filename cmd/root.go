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
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

var rootCmd = &cobra.Command{
	Use:     "lou",
	Version: "v0.2",
	Short:   chalk.Green.Color("Lou") + ", personal assitant at your service.",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` + chalk.Green.Color("Lou") + chalk.White.Color(", personal assitant at your service") + `.`,

	Example: `
` + chalk.Cyan.Color("lou") + ` help`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute
func Execute() {
	ε := rootCmd.Execute()
	if ε != nil {
		log.Fatal(ε)
		os.Exit(1)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// initialize config
func initializeConfig(κ *cobra.Command, configPath string, configName string) error {

	// initialize viper
	ω := viper.New()

	// collect config path & file from persistent flags
	ω.AddConfigPath(configPath)
	ω.SetConfigName(configName)

	// read config file
	ε := ω.ReadInConfig()
	if ε != nil {
		// okay if no config file
		_, ϙ := ε.(viper.ConfigFileNotFoundError)
		if !ϙ {
			// error if not parse config file
			return ε
		}
	}

	// bind flags viper
	bindFlags(κ, ω)

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// bind each cobra flag viper configuration
func bindFlags(κ *cobra.Command, ω *viper.Viper) {

	κ.Flags().VisitAll(func(ζ *pflag.Flag) {

		// apply viper config value flag
		if !ζ.Changed && ω.IsSet(ζ.Name) {
			ν := ω.Get(ζ.Name)
			κ.Flags().Set(ζ.Name, fmt.Sprintf("%v", ν))
		}
	})
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {

	// persistent flags
	rootCmd.PersistentFlags().StringP("location", "l", "/Users/drivas/Downloads/", "Location to clean")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
