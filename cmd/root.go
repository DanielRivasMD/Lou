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

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

////////////////////////////////////////////////////////////////////////////////////////////////////

var rootCmd = &cobra.Command{
	Use:   "Lou",
	Short: "Lou, personal assitant at your service",
	Long: `Daniel Rivas <danielrivasmd@gmail.com>

Lou, personal assitant at your service`,
	Version: "v0.1",
	Example: `
Lou clean
Lou biblo -r`,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	cobra.OnInitialize(initConfig)

	////////////////////////////////////////////////////////////////////////////////////////////////////

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file")

	////////////////////////////////////////////////////////////////////////////////////////////////////

}

////////////////////////////////////////////////////////////////////////////////////////////////////

func initConfig() {

	////////////////////////////////////////////////////////////////////////////////////////////////////

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".lou" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".lou")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////

}
