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
	"embed"
	"path/filepath"
	"sync"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

//go:embed docs.json
var docsFS embed.FS

////////////////////////////////////////////////////////////////////////////////////////////////////

const APP = "lou"
const VERSION = "v0.1.0"
const AUTHOR = "Daniel Rivas"
const EMAIL = "danielrivasmd@gmail.com"

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	onceRoot  sync.Once
	rootCmd   *cobra.Command
	rootFlags struct {
		verbose bool
		tabLayout string
		tabTarget string
	}
	configDirs configDir
)

type configDir struct {
	home   string
	lou    string
	layout string
	sh     string
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func InitDocs() {
	info := domovoi.AppInfo{
		Name:    APP,
		Version: VERSION,
		Author:  AUTHOR,
		Email:   EMAIL,
	}
	domovoi.SetGlobalDocsConfig(docsFS, info)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func GetRootCmd() *cobra.Command {
	onceRoot.Do(func() {
		d := horus.Must(domovoi.GlobalDocs())
		var err error
		rootCmd, err = d.MakeCmd("root", nil)
		horus.CheckErr(err)

		rootCmd.PersistentFlags().BoolVarP(&rootFlags.verbose, "verbose", "v", false, "Enable verbose diagnostics")
		rootCmd.Version = VERSION

		cobra.OnInitialize(initConfigDirs)
	})
	return rootCmd
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func Execute() {
	horus.CheckErr(GetRootCmd().Execute())
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func initConfigDirs() {
	var err error
	configDirs.home, err = domovoi.FindHome(rootFlags.verbose)
	horus.CheckErr(err, horus.WithCategory("init_error"), horus.WithMessage("getting home directory"))
	configDirs.lou = filepath.Join(configDirs.home, ".lou")
	configDirs.layout = filepath.Join(configDirs.lou, "layout")
	configDirs.sh = filepath.Join(configDirs.lou, "sh")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func BuildCommands() {
	root := GetRootCmd()
	root.AddCommand(
		CompletionCmd(),
		IdentityCmd(),

		CountCmd(),
		DiffCmd(),
		DocCmd(),
		DupCmd(),
		DxCmd(),
		GitCmd(),
		HiddenCmd(),
		KnitCmd(),
		PathCmd(),
		RallyCmd(),
		ShowCmd(),
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
