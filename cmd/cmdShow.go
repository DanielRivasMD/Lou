/*
Copyright © 2025 Daniel Rivas <danielrivasmd@gmail.com>

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
	"fmt"

	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var listShCmd = &cobra.Command{
	Use:     "show" + ` ` + chalk.Dim.TextStyle(chalk.Blue.Color("--file")) + ` ` + chalk.Dim.TextStyle("<file>"),
	Short:   "List shell functions from a specified file",
	Long:    helpShow,
	Example: exampleShow,

	Run: runShow,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	listFile string
)

type Function struct {
	Shell       string
	Name        string
	Description string
	Arguments   string
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(listShCmd)
	listShCmd.Flags().StringVarP(&listFile, "file", "f", "", "File to review")
	listShCmd.MarkFlagRequired("file")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runShow(cmd *cobra.Command, args []string) {

	// collect documentation
	functions, err := parseFile(listFile)
	horus.CheckErr(err)

	// generate & print Markdown
	markdown := generateMD(functions)
	fmt.Println(markdown)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
