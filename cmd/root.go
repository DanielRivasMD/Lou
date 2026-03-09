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
	"path/filepath"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

const APP = "lou"
const VERSION = "v0.1.0"
const NAME = "Daniel Rivas"
const EMAIL = "<danielrivasmd@gmail.com>"

////////////////////////////////////////////////////////////////////////////////////////////////////

var rootCmd = &cobra.Command{
	Use:     GetUse("root"),
	Long:    formatLongHelp(GetHelp("root")),
	Example: GetExample("root"),
	Version: VERSION,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func Execute() {
	horus.CheckErr(rootCmd.Execute())
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.PersistentFlags().BoolVarP(&rootFlags.verbose, "verbose", "v", false, "Enable verbose diagnostics")
	cobra.OnInitialize(initConfigDirs)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func initConfigDirs() {
	configDirs.home = func() string {
		h, e := domovoi.FindHome(rootFlags.verbose)
		horus.CheckErr(e, horus.WithCategory("init_error"), horus.WithMessage("getting home directory"))
		return h
	}()
	configDirs.lou = filepath.Join(configDirs.home, ".lou")
	configDirs.layout = filepath.Join(configDirs.lou, "layout")
	configDirs.sh = filepath.Join(configDirs.lou, "sh")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	configDirs configDir
	rootFlags  rootFlag
)

type configDir struct {
	home   string
	lou    string
	layout string
	sh     string
}

type rootFlag struct {
	verbose bool

	tabLayout string
	tabTarget string
}

////////////////////////////////////////////////////////////////////////////////////////////////////
