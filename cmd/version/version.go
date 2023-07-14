/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/

package version

import (
	"BitSrunLoginGo-CLI/cmd/template"
	"fmt"
	"github.com/spf13/cobra"
)

func NewVersionCmd(bitVer, cliVer, buildDate string) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "version",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(template.Version(bitVer, cliVer, buildDate))
		},
	}
	return cmd
}
