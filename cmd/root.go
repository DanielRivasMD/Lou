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

	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

// rootCmd
var rootCmd = &cobra.Command{
	Use:     "lou",
	Version: "v0.3",
	Short:   chalk.Green.Color("Lou") + ", personal assistant at your service.",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `

` + chalk.Green.Color("Lou") + `, personal assistant at your service
`,

	Example: `
` + chalk.Cyan.Color("lou") + ` ` + chalk.Magenta.Color("help") + `
`,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// Execute runs the root command.
func Execute() {
	err := rootCmd.Execute()
	horus.CheckErr(err)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// initializeConfig loads configuration from file into Viper.
func initializeConfig(cmd *cobra.Command, configPath, configName string) error {
	vip := viper.New()
	vip.AddConfigPath(configPath)
	vip.SetConfigName(configName)

	err := vip.ReadInConfig()
	if err != nil {
		_, notFound := err.(viper.ConfigFileNotFoundError)
		if !notFound {
			return err
		}
	}

	bindFlags(cmd, vip)
	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// bindFlags maps Viper values back into Cobra flags.
func bindFlags(cmd *cobra.Command, vip *viper.Viper) {
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if !flag.Changed && vip.IsSet(flag.Name) {
			val := vip.Get(flag.Name)
			cmd.Flags().Set(flag.Name, fmt.Sprintf("%v", val))
		}
	})
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior to main
func init() {
	// persistent flags
}

////////////////////////////////////////////////////////////////////////////////////////////////////
