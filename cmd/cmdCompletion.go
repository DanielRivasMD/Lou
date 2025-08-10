/*
Copyright © 2021 Daniel Rivas <danielrivasmd@gmail.com>

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

	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var completionCmd = &cobra.Command{
	Use:    "completion " + chalk.Dim.TextStyle(chalk.Italic.TextStyle("[bash|zsh|fish|powershell]")),
	Hidden: true,
	Short:  "Generate completion script for various shells.",
	Long:   helpCompletion,

	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),

	Run: runCompletion,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(completionCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var helpCompletion = fmt.Sprintf(`To load completions:

Bash:

  $ source <(%[1]s completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ %[1]s completion bash > /etc/bash_completion.d/%[1]s
  # macOS:
  $ %[1]s completion bash > $(brew --prefix)/etc/bash_completion.d/%[1]s

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it. You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ %[1]s completion zsh > "${fpath[1]}/_%[1]s"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ %[1]s completion fish | source

  # To load completions for each session, execute once:
  $ %[1]s completion fish > ~/.config/fish/completions/%[1]s.fish

PowerShell:

  PS> %[1]s completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> %[1]s completion powershell > %[1]s.ps1
  # and source this file from your PowerShell profile.
`, rootCmd.Name())

////////////////////////////////////////////////////////////////////////////////////////////////////

func runCompletion(cmd *cobra.Command, args []string) {
	// Generate the appropriate shell completion script based on the provided argument.
	switch args[0] {
	case "bash":
		horus.CheckErr(cmd.Root().GenBashCompletion(os.Stdout))
	case "zsh":
		horus.CheckErr(cmd.Root().GenZshCompletion(os.Stdout))
	case "fish":
		horus.CheckErr(cmd.Root().GenFishCompletion(os.Stdout, true))
	case "powershell":
		horus.CheckErr(cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout))
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
