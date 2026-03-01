/*
Copyright © 2026 Daniel Rivas <danielrivasmd@gmail.com>

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
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/DanielRivasMD/domovoi"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

//go:embed docs.json
var docsFS embed.FS

////////////////////////////////////////////////////////////////////////////////////////////////////

type CommandOpt func(*cobra.Command)

////////////////////////////////////////////////////////////////////////////////////////////////////

type DocEntry struct {
	Use                   string     `json:"use"`
	Aliases               []string   `json:"aliases,omitempty"`
	Hidden                bool       `json:"hidden,omitempty"`
	Short                 string     `json:"short,omitempty"`
	Long                  string     `json:"long"`
	ExampleUsages         [][]string `json:"example_usages,omitempty"`
	ValidArgs             []string   `json:"valid_args,omitempty"`
	DisableFlagsInUseLine bool       `json:"disable_flags_in_use_line,omitempty"`
}

////////////////////////////////////////////////////////////////////////////////////////////////////

type Docs struct {
	once    sync.Once
	entries map[string]DocEntry
	example map[string]string
	help    map[string]string
	short   map[string]string
	use     map[string]string
	hidden  map[string]bool
	loadErr error
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	globalDocs *Docs
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func formatHelp(text string, appName string) string {
	if text == "" {
		return ""
	}
	if strings.Contains(text, "%[1]s") || strings.Contains(text, "%s") {
		return fmt.Sprintf(text, appName)
	}
	return text
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func styleLongHelp(text string) string {
	if text == "" {
		return ""
	}
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "$") {
			lines[i] = chalk.White.Color(line)
		}
		if strings.HasPrefix(trimmed, "#") {
			lines[i] = chalk.Dim.TextStyle(chalk.Cyan.Color(line))
		}
	}
	return strings.Join(lines, "\n")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func styleDescription(text string) string {
	if text == "" {
		return ""
	}
	return chalk.Cyan.Color(chalk.Dim.TextStyle(text))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func authorHeader() string {
	return chalk.Bold.TextStyle(
		chalk.Green.Color(NAME),
	) + " " +
		chalk.Dim.TextStyle(
			chalk.Italic.TextStyle(EMAIL),
		) + "\n" +
		chalk.White.Color(APP) + " " +
		chalk.Bold.TextStyle(VERSION)
}

func formatLongHelp(help string) string {
	if help != "" {
		styledDesc := styleDescription(help)
		return authorHeader() + "\n\n" + styledDesc
	}
	return authorHeader()
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func (d *Docs) load() {
	d.entries = make(map[string]DocEntry)
	d.example = make(map[string]string)
	d.help = make(map[string]string)
	d.short = make(map[string]string)
	d.use = make(map[string]string)
	d.hidden = make(map[string]bool)

	data, err := docsFS.ReadFile("docs.json")
	if err != nil {
		d.loadErr = fmt.Errorf("failed to load embedded documentation: %w", err)
		return
	}

	if err := json.Unmarshal(data, &d.entries); err != nil {
		d.loadErr = fmt.Errorf("failed to parse embedded documentation: %w", err)
		return
	}

	for key, entry := range d.entries {
		d.use[key] = entry.Use
		d.short[key] = entry.Short
		d.hidden[key] = entry.Hidden

		formattedHelp := formatHelp(entry.Long, APP)
		d.help[key] = styleLongHelp(formattedHelp)

		if len(entry.ExampleUsages) > 0 {
			d.example[key] = domovoi.FormatExample(APP, entry.ExampleUsages...)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func (d *Docs) ensureLoaded() {
	d.once.Do(d.load)
	if d.loadErr != nil {
		panic(fmt.Sprintf("Documentation loading failed: %v", d.loadErr))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func (d *Docs) GetEntry(key string) (DocEntry, bool) {
	d.ensureLoaded()
	entry, exists := d.entries[key]
	return entry, exists
}

func (d *Docs) GetExample(key string) string {
	d.ensureLoaded()
	return d.example[key]
}

func (d *Docs) GetHelp(key string) string {
	d.ensureLoaded()
	return d.help[key]
}

func (d *Docs) GetHidden(key string) bool {
	d.ensureLoaded()
	return d.hidden[key]
}

func (d *Docs) GetShort(key string) string {
	d.ensureLoaded()
	return d.short[key]
}

func (d *Docs) GetUse(key string) string {
	d.ensureLoaded()
	return d.use[key]
}

func (d *Docs) GetAllEntries() map[string]DocEntry {
	d.ensureLoaded()
	return d.entries
}

func (d *Docs) ListCommands() []string {
	d.ensureLoaded()
	commands := make([]string, 0, len(d.entries))
	for k := range d.entries {
		commands = append(commands, k)
	}
	return commands
}

func (d *Docs) CommandExists(key string) bool {
	d.ensureLoaded()
	_, exists := d.entries[key]
	return exists
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func getGlobalDocs() *Docs {
	if globalDocs == nil {
		globalDocs = &Docs{}
	}
	return globalDocs
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func GetDocs() map[string]DocEntry {
	return getGlobalDocs().GetAllEntries()
}

func GetExample(key string) string {
	return getGlobalDocs().GetExample(key)
}

func GetHelp(key string) string {
	return getGlobalDocs().GetHelp(key)
}

func GetHidden(key string) bool {
	return getGlobalDocs().GetHidden(key)
}

func GetShort(key string) string {
	return getGlobalDocs().GetShort(key)
}

func GetUse(key string) string {
	return getGlobalDocs().GetUse(key)
}

func GetDocEntry(key string) (DocEntry, bool) {
	return getGlobalDocs().GetEntry(key)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func MakeCmd(key string, run func(*cobra.Command, []string), opts ...CommandOpt) *cobra.Command {
	d := getGlobalDocs()

	entry, exists := d.GetEntry(key)
	if !exists {
		keys := d.ListCommands()
		log.Fatalf("No documentation found for command: %s. Available keys: %v", key, keys)
	}

	cmd := &cobra.Command{
		Use:     d.GetUse(key),
		Short:   d.GetShort(key),
		Long:    formatLongHelp(d.GetHelp(key)),
		Example: d.GetExample(key),
		Aliases: entry.Aliases,
		Hidden:  d.GetHidden(key),
		Run:     run,
	}

	if len(entry.ValidArgs) > 0 {
		cmd.ValidArgs = entry.ValidArgs
	}

	if entry.DisableFlagsInUseLine {
		cmd.DisableFlagsInUseLine = true
	}

	for _, opt := range opts {
		opt(cmd)
	}

	return cmd
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func WithArgs(validator cobra.PositionalArgs) CommandOpt {
	return func(cmd *cobra.Command) {
		cmd.Args = validator
	}
}

func WithValidArgs(args []string) CommandOpt {
	return func(cmd *cobra.Command) {
		cmd.ValidArgs = args
	}
}

func WithAliases(aliases []string) CommandOpt {
	return func(cmd *cobra.Command) {
		cmd.Aliases = aliases
	}
}

func WithDisableFlagParsing(disable bool) CommandOpt {
	return func(cmd *cobra.Command) {
		cmd.DisableFlagParsing = disable
	}
}

func WithPreRun(preRun func(*cobra.Command, []string)) CommandOpt {
	return func(cmd *cobra.Command) {
		cmd.PreRun = preRun
	}
}

func WithPostRun(postRun func(*cobra.Command, []string)) CommandOpt {
	return func(cmd *cobra.Command) {
		cmd.PostRun = postRun
	}
}

func WithSilenceErrors(silence bool) CommandOpt {
	return func(cmd *cobra.Command) {
		cmd.SilenceErrors = silence
	}
}

func WithSilenceUsage(silence bool) CommandOpt {
	return func(cmd *cobra.Command) {
		cmd.SilenceUsage = silence
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
