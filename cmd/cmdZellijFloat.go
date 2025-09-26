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

var zellijFloatCmd = &cobra.Command{}

var batCmd = &cobra.Command{
	Use:     "bat " + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<file>")),
	Short:   "View data in a floating zellij pane using bat",
	Long:    helpBat,
	Example: exampleBat,

	Run: runBat,
}

var brootCmd = &cobra.Command{
	Use:     "broot",
	Aliases: []string{"br"},
	Short:   "Browse files in a floating zellij pane using broot",

	Run: runBroot,
}

var ezaCmd = &cobra.Command{
	Use:     "eza " + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<path>")),
	Short:   "View data in a floating zellij pane using eza",
	Long:    helpEza,
	Example: exampleEza,

	Run: runEza,
}

var floatCmd = &cobra.Command{
	Use:     "float",
	Short:   "Launch floating zellij pane with ease",
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
	Short:   "Launch lazygit in a floating zellij pane",
	Long:    helpLazygit,
	Example: exampleLazygit,

	Run: runLazygit,
}

var mdcatCmd = &cobra.Command{
	Use:   "mdcat",
	Short: "Render file with mdcat in a floating zellij pane",

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
	Short:   "Anchor & resize a zellij floating pane",
	Long:    helpResize,
	Example: exampleResize,

	Args:      cobra.MaximumNArgs(1),
	ValidArgs: validLayouts,

	Run: runResize,
}

var watchCmd = &cobra.Command{
	Use:     "watch",
	Short:   "Launch watcher on a floating zellij pane with ease",
	Long:    helpWatch,
	Example: exampleZFWatch,

	Run: runWatch,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

type zellijOpt func(*zellijFloat)

type zellijFloat struct {
	name        string
	closeOnExit bool
	pinned      bool
	layout      string
	command     string
	args        []string
	customGeom  *Geometry
}

func newZellijFloat(opts ...zellijOpt) zellijFloat {
	zf := zellijFloat{
		closeOnExit: false,
		pinned:      false,
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

func withGeometry(g Geometry) zellijOpt {
	return func(z *zellijFloat) {
		z.layout = "__custom__"
		z.customGeom = &g
	}
}

func (zl zellijFloat) Cmd() string {
	var geom Geometry
	if zl.layout == "__custom__" && zl.customGeom != nil {
		geom = *zl.customGeom
	} else {
		geom, _ = resolveLayoutGeometry(zl.layout, flagG)
	}

	flags := []string{"--name " + zl.name}
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

	cmd := fmt.Sprintf("zellij run --floating %s -- %s", strings.Join(flags, " "), zl.command)
	if len(zl.args) > 0 {
		cmd += " " + strings.Join(zl.args, " ")
	}
	return cmd
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var flagG Geometry

type Geometry struct {
	width  string
	height string
	x      string
	y      string
}

// TODO: redundant defaults?
func (g Geometry) OverrideWith(flags Geometry) Geometry {
	return Geometry{
		height: override(g.height, flags.height, "100%"),
		width:  override(g.width, flags.width, "95%"),
		x:      override(g.x, flags.x, "10"),
		y:      override(g.y, flags.y, "0"),
	}
}

func geometryFromFlags() Geometry {
	return Geometry{
		height: flagG.height,
		width:  flagG.width,
		x:      flagG.x,
		y:      flagG.y,
	}
}

///////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(zellijFloatCmd)
	rootCmd.AddCommand(batCmd, brootCmd, ezaCmd, floatCmd, helixCmd, lazygitCmd, mdcatCmd, microCmd, resizeCmd, watchCmd)
	zellijFloatCmd.AddCommand(batCmd, brootCmd, ezaCmd, floatCmd, helixCmd, lazygitCmd, mdcatCmd, microCmd, resizeCmd, watchCmd)

	rootCmd.PersistentFlags().StringVarP(&flagG.height, "height", "H", "100%", "pane height")
	rootCmd.PersistentFlags().StringVarP(&flagG.width, "width", "W", "95%", "pane width")
	rootCmd.PersistentFlags().StringVarP(&flagG.x, "x", "X", "10", "horizontal offset")
	rootCmd.PersistentFlags().StringVarP(&flagG.y, "y", "Y", "0", "vertical offset")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runBat(cmd *cobra.Command, args []string) {
	op := "lou.zellij.bat"

	batArgs := []string{"--paging=always"}
	if len(args) > 0 {
		batArgs = append(batArgs, args[0])
	}

	zl := newZellijFloat(
		withName("bat"),
		withGeometry(geometryFromFlags()),
		withCommand("bat"),
		withArgs(batArgs...),
		withCloseOnExit(true),
	)

	horus.CheckErr(
		domovoi.ExecSh(zl.Cmd()),
		horus.WithOp(op),
		horus.WithCategory("shell_command"),
		horus.WithMessage("Failed to launch bat"),
		horus.WithDetails(map[string]any{
			"command": zl.Cmd(),
		}),
	)
}

func runBroot(cmd *cobra.Command, args []string) {
	op := "lou.zellij.broot"

	zl := newZellijFloat(
		withName("broot"),
		withGeometry(geometryFromFlags()),
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
			withGeometry(geometryFromFlags()),
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

	ezaArgs := []string{"--header", "--long", "--icons", "--classify", "--git", "--group", "--color=always"}
	if len(args) > 0 {
		ezaArgs = append(ezaArgs, args[0])
	}

	opts := []zellijOpt{
		withName("eza"),
		withGeometry(geometryFromFlags()),
		withCommand("eza"),
		withArgs(ezaArgs...),
	}

	zl := newZellijFloat(opts...)

	horus.CheckErr(
		domovoi.ExecSh(zl.Cmd()),
		horus.WithOp(op),
		horus.WithCategory("shell_command"),
		horus.WithMessage("Failed to launch eza"),
		horus.WithDetails(map[string]any{"command": zl.Cmd()}),
	)
}

func runFloat(cmd *cobra.Command, args []string) {
	op := "lou.zellij.float"

	floatLayout := "default"
	if len(args) == 1 {
		floatLayout = args[0]
	}

	geom, err := resolveWithFlags(floatLayout)
	horus.CheckErr(err)

	zl := newZellijFloat(
		withName("canvas"),
		withGeometry(geom),
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

	floatLayout := "full"
	if len(args) == 1 {
		floatLayout = args[0]
	}

	geom, err := resolveWithFlags(floatLayout)
	horus.CheckErr(err)

	zl := newZellijFloat(
		withName("lazygit"),
		withGeometry(geom),
		withCommand("lazygit"),
		withCloseOnExit(true),
		withPinned(true),
	)

	horus.CheckErr(
		domovoi.ExecSh(zl.Cmd()),
		horus.WithOp(op),
		horus.WithCategory("shell_command"),
		horus.WithMessage("Failed to launch lazygit"),
		horus.WithDetails(map[string]any{"command": zl.Cmd()}),
	)
}

func runMDcat(cmd *cobra.Command, args []string) {
	op := "lou.zellij.mdcat"

	// TODO: add one-liner error
	if len(args) < 1 {
		horus.CheckErr(
			horus.NewHerrorErrorf(op, "mdcat command requires a file argument"),
		)
	}
	file := args[0]

	zl := newZellijFloat(
		withName("canvas"),
		withGeometry(geometryFromFlags()),
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

	resizeLayout := "default"
	if len(args) == 1 {
		resizeLayout = args[0]
	}

	geom, err := resolveWithFlags(resizeLayout)
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
		horus.WithDetails(map[string]any{
			"command": cmdResize,
		}))
}

func runWatch(cmd *cobra.Command, args []string) {
	op := "lou.zellij.watch"

	// TODO: check if `just watch` exist
	zl := newZellijFloat(
		withName("watch"),
		withGeometry(geometryFromFlags()),
		withCommand("just"),
		withArgs("watch"),
		withCloseOnExit(true),
		// withFloating(true),
		withPinned(true),
	)

	horus.CheckErr(
		domovoi.ExecSh(zl.Cmd()),
		horus.WithOp(op),
		horus.WithCategory("shell_command"),
		horus.WithMessage("Failed to launch watch command"),
		horus.WithDetails(map[string]any{"command": zl.Cmd()}),
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func override(preset, flag, fallback string) string {
	if preset != "" {
		return preset
	}
	if flag != fallback {
		return flag
	}
	return fallback
}

var layoutPresets = map[string]Geometry{
	"full":         {"100%", "100%", "0", "0"},
	"half-left":    {"50%", "100%", "0", "0"},
	"half-right":   {"50%", "100%", "50%", "0"},
	"top-left":     {"50%", "50%", "0", "0"},
	"bottom-left":  {"50%", "53%", "0", "52%"},
	"top-right":    {"50%", "50%", "50%", "0"},
	"bottom-right": {"50%", "53%", "50%", "52%"},
	"default":      {"95%", "100%", "10", "0"},
}

var validLayouts = func() []string {
	keys := make([]string, 0, len(layoutPresets))
	for k := range layoutPresets {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}()

func resolveLayoutGeometry(name string, flags Geometry) (Geometry, error) {
	preset, ok := layoutPresets[name]
	if !ok {
		return Geometry{}, fmt.Errorf("unknown layout %q (must be one of %v)", name, validLayouts)
	}
	return preset.OverrideWith(flags), nil
}

func resolveWithFlags(name string) (Geometry, error) {
	return resolveLayoutGeometry(name, flagG)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
