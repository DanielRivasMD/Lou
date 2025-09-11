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
	"sort"
	"strings"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// BUG: double check running errors
var zellijFloatCmd = &cobra.Command{}

var batCmd = &cobra.Command{
	Use:     "bat " + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<file>")),
	Short:   "View data in a floating zellij window using bat",
	Long:    helpBat,
	Example: exampleBat,

	Run: runBat,
}

var brootCmd = &cobra.Command{
	Use:     "broot",
	Aliases: []string{"br"},
	Short:   "browse files in a floating zellij window using broot",

	Run: runBroot,
}

var ezaCmd = &cobra.Command{
	Use:     "eza " + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<path>")),
	Short:   "View data in a floating zellij window using eza",
	Long:    helpEza,
	Example: exampleEza,

	Run: runEza,
}

var floatCmd = &cobra.Command{
	Use:     "float",
	Short:   "launch floating zellij window with ease",
	Long:    helpFloat,
	Example: exampleFloat,

	Args:      cobra.MaximumNArgs(1),
	ValidArgs: validLayouts,

	Run: runFloat,
}

var helixCmd = &cobra.Command{
	Use:     "helix " + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<file>")),
	Aliases: []string{"hx"},
	Short:   helpEditorShort("helix"),
	Long:    helpEditorLong("helix"),
	// Example: helpEditorExample("helix"),

	Run: runEditor("hx"),
}

var lazygitCmd = &cobra.Command{
	Use:     "lazygit",
	Aliases: []string{"lg"},
	Short:   "lazygit in a floating zellij window",
	Long:    helpZFLazygit,
	Example: exampleZFLazygit,

	Run: runLazygit,
}

var mdcatCmd = &cobra.Command{
	Use:   "mdcat",
	Short: "",

	Run: runMDcat,
}

var microCmd = &cobra.Command{
	Use:     "micro " + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<file>")),
	Aliases: []string{"mc"},
	Short:   helpEditorShort("micro"),
	Long:    helpEditorLong("micro"),
	// Example: helpEditorExample("micro"),

	Run: runEditor("micro"),
}

var resizeCmd = &cobra.Command{
	Use:     "resize",
	Short:   "anchor and resize a random floating pane",
	Long:    helpZFResize,
	Example: exampleZFResize,

	Args:      cobra.MaximumNArgs(1),
	ValidArgs: validLayouts,

	Run: runResize,
}

var watchCmd = &cobra.Command{
	Use:     "watch",
	Short:   "launch watcher on a floating zellij window with ease",
	Long:    helpZFWatch,
	Example: exampleZFWatch,

	Run: runWatch,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

type zellijOpt func(*zellijFloat)

type zellijFloat struct {
	name        string
	closeOnExit bool
	floating    bool
	pinned      bool
	layout      string
	command     string
	args        []string
}

func newZellijFloat(opts ...zellijOpt) zellijFloat {
	zf := zellijFloat{
		floating:    true,
		pinned:      true,
		closeOnExit: false,
		layout:      "default",
	}
	for _, opt := range opts {
		opt(&zf)
	}
	return zf
}

func withName(name string) zellijOpt {
	return func(z *zellijFloat) {
		z.name = name
	}
}

func withLayout(layout string) zellijOpt {
	return func(z *zellijFloat) {
		z.layout = layout
	}
}

func withCommand(cmd string) zellijOpt {
	return func(z *zellijFloat) {
		z.command = cmd
	}
}

func withArgs(args ...string) zellijOpt {
	return func(z *zellijFloat) {
		z.args = args
	}
}

func withFloating(v bool) zellijOpt {
	return func(z *zellijFloat) {
		z.floating = v
	}
}

func withPinned(v bool) zellijOpt {
	return func(z *zellijFloat) {
		z.pinned = v
	}
}

func withCloseOnExit(v bool) zellijOpt {
	return func(z *zellijFloat) {
		z.closeOnExit = v
	}
}

func (zl zellijFloat) Cmd() string {
	geom, _ := resolveLayoutGeometry(zl.layout)

	flags := []string{"--name " + zl.name}
	if zl.floating {
		flags = append(flags, "--floating")
	}
	if zl.pinned {
		flags = append(flags, "--pinned true")
	}
	if zl.closeOnExit {
		flags = append(flags, "--close-on-exit")
	}

	flags = append(flags,
		"--height "+geom.height,
		"--width "+geom.width,
		"--x "+geom.x,
		"--y "+geom.y,
	)

	cmd := fmt.Sprintf("zellij run %s -- %s", strings.Join(flags, " "), zl.command)
	if len(zl.args) > 0 {
		cmd += " " + strings.Join(zl.args, " ")
	}
	return cmd
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	dLayout = newFloatCoor()
	flags   zellijFlags
)

type zellijFlags struct {
	layout floatCoor
}

type floatCoor struct {
	height string
	width  string
	x      string
	y      string
}

func newFloatCoor() floatCoor {
	return floatCoor{
		height: "100%",
		width:  "95%",
		x:      "10",
		y:      "0",
	}
}

func (fc floatCoor) String() string {
	return fmt.Sprintf(
		`--height %s \
--width %s \
--x %s \
--y %s \`,
		fc.height,
		fc.width,
		fc.x,
		fc.y,
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(zellijFloatCmd)
	rootCmd.AddCommand(batCmd, brootCmd, ezaCmd, floatCmd, helixCmd, lazygitCmd, mdcatCmd, microCmd, resizeCmd, watchCmd)
	zellijFloatCmd.AddCommand(batCmd, brootCmd, ezaCmd, floatCmd, helixCmd, lazygitCmd, mdcatCmd, microCmd, resizeCmd, watchCmd)

	zellijFloatCmd.PersistentFlags().StringVarP(&flags.layout.height, "height", "H", dLayout.height, "pane height as percentage")
	zellijFloatCmd.PersistentFlags().StringVarP(&flags.layout.width, "width", "W", dLayout.width, "pane width as percentage")
	zellijFloatCmd.PersistentFlags().StringVarP(&flags.layout.x, "x", "X", dLayout.x, "horizontal offset as percentage")
	zellijFloatCmd.PersistentFlags().StringVarP(&flags.layout.y, "y", "Y", dLayout.y, "vertical offset as percentage")

	layoutPresets["default"] = LayoutGeometry{
		width:  dLayout.width,
		height: dLayout.height,
		x:      dLayout.x,
		y:      dLayout.y,
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runBat(cmd *cobra.Command, args []string) {
	op := "lou.zellij.bat"
	zl := newZellijFloat(
		withName("bat"),
		withLayout(flags.layout.String()),
		withCommand("eza"),
		withArgs("--paging=always"),
	)

	if len(args) > 0 {
		zl.args = append(zl.args, args[0])
	}

	horus.CheckErr(
		domovoi.ExecSh(zl.Cmd()),
		horus.WithOp(op),
		horus.WithCategory("shell_command"),
		horus.WithMessage("Failed to execute mbombo forge command"),
		horus.WithDetails(map[string]any{
			"command": zl.Cmd(),
		}),
	)
}

func runBroot(cmd *cobra.Command, args []string) {
	op := "lou.zellij.broot"

	zl := newZellijFloat(
		withName("broot"),
		withLayout("default"),
		withCommand("broot"),
		withArgs("--dates", "--sizes", "--permissions", "--hidden", "--git-ignored", "--show-git-info", "--sort-by-type-dirs-first"),
		withCloseOnExit(true),
	)

	horus.CheckErr(
		domovoi.ExecSh(zl.Cmd()),
		horus.WithOp(op),
		horus.WithCategory("shell_command"),
		horus.WithMessage("Failed to launch broot"),
		horus.WithDetails(map[string]any{"command": zl.Cmd()}),
	)
}

func runEditor(call string, editorOverride ...string) func(cmd *cobra.Command, args []string) {
	editor := call
	if len(editorOverride) > 0 {
		editor = editorOverride[0]
	}

	return func(cmd *cobra.Command, args []string) {
		op := "lou.zellij.editor"

		opts := []zellijOpt{
			withName(editor),
			withLayout(flags.layout.String()),
			withCommand(call),
			withCloseOnExit(true),
		}
		if len(args) > 0 {
			opts = append(opts, withArgs(args[0]))
		}

		zl := newZellijFloat(opts...)

		horus.CheckErr(
			domovoi.ExecSh(zl.Cmd()),
			horus.WithOp(op),
			horus.WithCategory("shell_command"),
			horus.WithMessage("Failed to launch editor"),
			horus.WithDetails(map[string]any{"command": zl.Cmd()}),
		)
	}
}

func runEza(cmd *cobra.Command, args []string) {
	op := "lou.zellij.eza"
	zl := newZellijFloat(
		withName("eza"),
		withLayout(flags.layout.String()),
		withCommand("eza"),
		withArgs("--header", "--long", "--icons", "--classify", "--git", "--group", "--color=always"),
	)

	if len(args) > 0 {
		zl.args = append(zl.args, args[0])
	}

	horus.CheckErr(
		domovoi.ExecSh(zl.Cmd()),
		horus.WithOp(op),
		horus.WithCategory("shell_command"),
		horus.WithMessage("Failed to execute mbombo forge command"),
		horus.WithDetails(map[string]any{
			"command": zl.Cmd(),
		}),
	)
}

func runFloat(cmd *cobra.Command, args []string) {
	op := "lou.zellij.float"

	layout := "default"
	if len(args) == 1 {
		layout = args[0]
	}

	zl := newZellijFloat(
		withName("canvas"),
		withLayout(layout),
		withCommand("zsh"),
		withCloseOnExit(true),
	)

	horus.CheckErr(
		domovoi.ExecSh(zl.Cmd()),
		horus.WithOp(op),
		horus.WithCategory("shell_command"),
		horus.WithMessage("Failed to launch floating shell"),
		horus.WithDetails(map[string]any{"command": zl.Cmd()}),
	)
}

func runLazygit(cmd *cobra.Command, args []string) {
	op := "lou.zellij.lazygit"
	zl := newZellijFloat(
		withName("lazygit"),
		withLayout(flags.layout.String()),
		withCommand("lazygit"),
	)

	horus.CheckErr(
		domovoi.ExecSh(zl.Cmd()),
		horus.WithOp(op),
		horus.WithCategory("shell_command"),
		horus.WithMessage("Failed to execute mbombo forge command"),
		horus.WithDetails(map[string]any{
			"command": zl.Cmd(),
		}),
	)
}

func runMDcat(cmd *cobra.Command, args []string) {
	op := "lou.zellij.mdcat"

	// TODO: add one line error message
	if len(args) < 1 {
		horus.CheckErr(
			horus.NewHerrorErrorf(op, "mdcat command requires a file argument"),
		)
	}
	file := args[0]

	zl := newZellijFloat(
		withName("canvas"),
		withLayout(flags.layout.String()),
		withCommand("mdcat"),
		withArgs("--paginate", file),
		withCloseOnExit(true),
	)

	horus.CheckErr(
		domovoi.ExecSh(zl.Cmd()),
		horus.WithOp(op),
		horus.WithCategory("shell_command"),
		horus.WithMessage("Failed to launch mdcat"),
		horus.WithDetails(map[string]any{"command": zl.Cmd()}),
	)
}

func runResize(cmd *cobra.Command, args []string) {
	op := "lou.zellij.resize"

	layout := "default"
	if len(args) == 1 {
		layout = args[0]
	}

	geom, err := resolveLayoutGeometry(layout)
	horus.CheckErr(err)

	cmdResize := fmt.Sprintf(`
zellij action rename-pane float
zellij action change-floating-pane-coordinates --pane-id $ZELLIJ_PANE_ID \
--height %s \
--width %s \
--x %s \
--y %s`, geom.height, geom.width, geom.x, geom.y)

	horus.CheckErr(
		domovoi.ExecSh(cmdResize),
		horus.WithOp(op),
		horus.WithCategory("shell_command"),
		horus.WithMessage("Failed to resize floating pane"),
		horus.WithDetails(map[string]any{"command": cmdResize}),
	)
}

// TODO: test whether `just watch` exists
func runWatch(cmd *cobra.Command, args []string) {
	op := "lou.zellij.watch"

	zl := newZellijFloat(
		withName("watch"),
		withLayout(flags.layout.String()),
		withCommand("just"),
		withArgs("watch"),
		withCloseOnExit(true),
		withFloating(true),
		withPinned(true),
	)

	horus.CheckErr(
		domovoi.ExecSh(zl.Cmd()),
		horus.WithOp(op),
		horus.WithCategory("shell_command"),
		horus.WithMessage("Failed to launch watch command"),
		horus.WithDetails(map[string]any{
			"command": zl.Cmd(),
		}),
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// define type to hold parameters
type LayoutGeometry struct {
	width, height, x, y string
}

// declare a map from layout name → geometry
var layoutPresets = map[string]LayoutGeometry{
	"full": {
		width:  "100%",
		height: "100%",
		x:      "0",
		y:      "0",
	},
	"half-left": {
		width:  "50%",
		height: "100%",
		x:      "0",
		y:      "0",
	},
	"half-right": {
		width:  "50%",
		height: "100%",
		x:      "50%",
		y:      "0",
	},
	"top-left": {
		width:  "45%",
		height: "45%",
		x:      "0",
		y:      "0",
	},
	"bottom-left": {
		width:  "45%",
		height: "45%",
		x:      "0",
		y:      "60%",
	},
	"top-right": {
		width:  "45%",
		height: "45%",
		x:      "60%",
		y:      "0",
	},
	"bottom-right": {
		width:  "45%",
		height: "45%",
		x:      "60%",
		y:      "60%",
	},
}

// derive validLayouts slice from map keys
var validLayouts = func() []string {
	keys := make([]string, 0, len(layoutPresets))
	for name := range layoutPresets {
		keys = append(keys, name)
	}
	sort.Strings(keys)
	return keys
}()

////////////////////////////////////////////////////////////////////////////////////////////////////

// pickValue applies bottom -> top override:
//   - start with defaultVal
//   - if presetVal is non-empty, use
//   - if flagVal differs from defaultVal, use
func pickValue(defaultVal, presetVal, flagVal string) string {
	result := defaultVal
	if presetVal != "" {
		result = presetVal
	}
	if flagVal != defaultVal {
		result = flagVal
	}
	return result
}

// look up named preset, apply overrides from flags, and return final LayoutGeometry
func resolveLayoutGeometry(layoutName string) (LayoutGeometry, error) {
	geom, ok := layoutPresets[layoutName]
	if !ok {
		return LayoutGeometry{}, fmt.Errorf(
			"unknown layout %q (must be one of %v)",
			layoutName, validLayouts,
		)
	}

	// override using coorFlags
	geom.width = pickValue("95%", geom.width, flags.layout.width)
	geom.height = pickValue("100%", geom.height, flags.layout.height)
	geom.x = pickValue("10", geom.x, flags.layout.x)
	geom.y = pickValue("0", geom.y, flags.layout.y)

	return geom, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
