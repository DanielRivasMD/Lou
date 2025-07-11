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

import (
	"github.com/DanielRivasMD/domovoi"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

// diagnosticsCmd
var diagnosticsCmd = &cobra.Command{
	Use:   "diagnostics",
	Short: "" + chalk.Yellow.Color("") + ".",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `
`,

	Example: `
` + chalk.Cyan.Color("") + ` help ` + chalk.Yellow.Color("") + chalk.Yellow.Color("shelldiagnostics"),

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {

		// base command
		cmdDiagnostics := `
#────────────────────────────────────────────────────────────
# Terminal diagnostics on demand
#────────────────────────────────────────────────────────────
echo -e "[ ZSH SESSION DIAGNOSTICS ]"
echo "───────────────────────────────"

echo -n "ZSH Version: "; zsh --version

echo -n "Loaded Profile: "; [[ -f "$HOME/.profile" ]] && echo "✓" || echo "⚠ Not Found"
echo -n "Sheldon Plugin Manager: "; command -v sheldon >/dev/null && echo "✓" || echo "⚠ Missing"
echo -n "Starship Prompt: "; command -v starship >/dev/null && echo "✓" || echo "⚠ Missing"
echo -n "Atuin History Manager: "; command -v atuin >/dev/null && echo "✓" || echo "⚠ Missing"

echo -n "Zoxide: "; command -v zoxide >/dev/null && echo "✓" || echo "⚠ Missing"
echo -n "Yazi: "; command -v yazi >/dev/null && echo "✓" || echo "⚠ Missing"
echo -n "fzf: "; command -v fzf >/dev/null && echo "✓" || echo "⚠ Missing"
echo ""

echo "[ Environment Variables ]"
env | grep -E 'ARCHIVE|EX_SITU|IN_SITU|IN_SILICO|ZSH_COMPLETION|BARTIB_FILE|VISUAL|EDITOR|PAGER|BAT_PAGER|RCOLUMNS|MANWIDTH|STARSHIP_CONFIG|ZELLIJ_CONFIG_DIR|ZELLIJ_CONFIG_FILE|ATUIN_NOBIND|GOPATH|GOROOT|LANG|LC_ALL|HISTFILE|ZDOTDIR' | sort
echo "───────────────────────────────"
`

		// execute command
		domovoi.ExecCmd("zsh", "-c", cmdDiagnostics)
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(diagnosticsCmd)

	// flags
}

////////////////////////////////////////////////////////////////////////////////////////////////////
