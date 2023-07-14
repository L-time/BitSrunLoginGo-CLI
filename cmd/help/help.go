package help

import "github.com/spf13/cobra"

func NewHelpCmd(root *cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "help <command>",
		Short:  "获取帮助",
		Long:   `获取某个命令的详细帮助说明`,
		Run:    root.HelpFunc(),
		Hidden: false,
	}

	return cmd
}
