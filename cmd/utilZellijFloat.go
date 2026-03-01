///////////////////////////////////////////////////////////////////////////////////////////////////

package cmd

///////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

///////////////////////////////////////////////////////////////////////////////////////////////////

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

///////////////////////////////////////////////////////////////////////////////////////////////////

func createZellijCommand(
	key string,
	run func(*cobra.Command, []string),
	parent *cobra.Command,
	opts ...CommandOpt,
) (*cobra.Command, *cobra.Command) {
	hierCmd := MakeCmd("zellij-"+key, run, opts...)
	topCmd := MakeCmd(key, run, opts...)

	parent.PersistentFlags().VisitAll(func(flag *pflag.Flag) {
		topCmd.Flags().AddFlag(flag)
	})
	parent.AddCommand(hierCmd)

	return hierCmd, topCmd
}

///////////////////////////////////////////////////////////////////////////////////////////////////
