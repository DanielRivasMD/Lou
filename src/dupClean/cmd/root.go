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
  "log"

  "lineBreaks"
  "shellCall"

  "github.com/labstack/gommon/color"

  "github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
  Use:   "dupClean",
  Short: "dupClean purges downloaded duplicates",
  Long: `dupClean purges downloaded duplicates`,

  // Uncomment the following line if your bare application
  // has an action associated with it:
  Run: func(cmd *cobra.Command, args []string) {

    // lineBreaks
    lineBreaks.LineBreaks()

    // shellCall
    err, out, errout := shellCall.ShellCall("/Users/drivas/Factorem/Lou/src/dupClean/shell/dupClean.sh")
    if err != nil {
      log.Printf("error: %v\n", err)
    }

    // stdout
    color.Println(color.Cyan(out, color.B))

    // stderr
    if errout != "" {
      color.Println(color.Red(errout, color.B))
    }

    // lineBreaks
    lineBreaks.LineBreaks()

  },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
  fmt.Println(err)
  os.Exit(1)
  }
}
