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
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

// bluetoothCmd
var bluetoothCmd = &cobra.Command{
	Use:   "bluetooth",
	Short: "" + chalk.Yellow.Color("") + ".",
	Long: chalk.Green.Color(chalk.Bold.TextStyle("Daniel Rivas ")) + chalk.Dim.TextStyle(chalk.Italic.TextStyle("<danielrivasmd@gmail.com>")) + `
`,

	Example: `
` + chalk.Cyan.Color("") + ` help ` + chalk.Yellow.Color("") + chalk.Yellow.Color("bluetooth"),

	////////////////////////////////////////////////////////////////////////////////////////////////////

}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(bluetoothCmd)

	// flags
}

////////////////////////////////////////////////////////////////////////////////////////////////////


// pairedDevices() {
// osascript << EOF
//     use framework "IOBluetooth"
//     use scripting additions
//     set _results to {}

//     repeat with device in (current application's IOBluetoothDevice's pairedDevices() as list)
//         if device's isPaired()
//             set _address to device's addressString as string
//             set _name to device's nameOrAddress as string
//             set _isConnected to device's isConnected as string
//             if _isConnected = "1"
//                 set _isConnected to "✔"
//             else 
//                 set _isConnected to "✗"
//             end if
//             set end of _results to {_address, "\t", _name, "\t", _isConnected, "\n"}
//         end if
//     end repeat

//     return _results as string
// EOF
// }

// connect() {
// local address=$1
// osascript << EOF
//     use framework "IOBluetooth"
//     use scripting additions

//     repeat with device in (current application's IOBluetoothDevice's pairedDevices() as list)
//         set _address to device's addressString() as string
//         if _address = "${address}"
//             if device's isConnected()
//                 device's closeConnection()
//             else
//                 device's openConnection()
//             end if
//         end if
//     end repeat
// EOF
// }

// main() {
//     local selected=("$(pairedDevices \
//         | sed '/^$/d' \
//         | fzf \
//             --delimiter $'\t' --with-nth 2,3 \
//             --preview \
//                 'system_profiler SPBluetoothDataType -json 2>/dev/null \
//                     | jq -r ".SPBluetoothDataType[].device_title[]["\"{2}\""]
//                     | select(type != \"null\")"' \
//     )")
//     echo "${selected[@]}" | while read line; do
//         local address name
//         read address name <<< $(echo "$line" | cut -f1-)
//         echo "${name}"
//         connect ${address} >/dev/null
//     done
// }

////////////////////////////////////////////////////////////////////////////////////////////////////
