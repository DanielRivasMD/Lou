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

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var rootCmd = &cobra.Command{
	Use:     "lou",
	Long:    helpRoot,
	Example: exampleRoot,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func Execute() {
	horus.CheckErr(rootCmd.Execute())
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

// WIP: add cobra.OnInitialize to declare paths
func init() {
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var helpRoot = chalk.Bold.TextStyle(chalk.Green.Color("Daniel Rivas ")) +
	chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) +
	chalk.Dim.TextStyle(chalk.Cyan.Color("\n\npersonal assistant at your service"))

var exampleRoot = chalk.White.Color("lou") + " " + chalk.Bold.TextStyle(chalk.White.Color("help"))

////////////////////////////////////////////////////////////////////////////////////////////////////
