/*
Copyright © 2023 Leave_Time

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package root

import (
	"BitSrunLoginGo-CLI/cmd/config"
	"BitSrunLoginGo-CLI/cmd/version"
	"github.com/spf13/cobra"
)

func NewRootCmd(bitVer, cliVer, buildDate string) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "bit-cli <command> <subcommand> [flags]",
		Short: "BitSrunLoginGo CLI",
		Long:  `BitSrunLoginGo的CLI实现，为深澜网络认证提供终端式自助操作`,
	}
	cmd.CompletionOptions.DisableDefaultCmd = true

	cmd.AddCommand(version.NewVersionCmd(bitVer, cliVer, buildDate))
	cmd.AddCommand(config.NewConfigCmd())
	//cmd.SetHelpCommand(help.NewHelpCmd(cmd))

	//cmd.SetUsageTemplate(template.Usage())
	//cmd.SetHelpTemplate(template.Help())

	cmd.PersistentFlags().BoolP("help", "h", false, "获取帮助")

	return cmd
}
