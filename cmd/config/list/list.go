package list

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/user"
	"path/filepath"
)

func NewConfigListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "列出配置",
		Long:    `列出所有的配置文件`,
		Aliases: []string{"ls", "all"},
		RunE:    list,
	}
	return cmd
}

func list(cmd *cobra.Command, args []string) error {
	var files []string
	path, err := user.Current()
	if err != nil {
		return err
	}
	roots := path.HomeDir + "/BitSrunLogin/"
	err = filepath.Walk(roots, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(info.Name()) == "" {
			return nil
		}
		filename := info.Name()
		extension := filepath.Ext(filename)
		files = append(files)
		name := filename[0 : len(filename)-len(extension)]
		files = append(files, name)
		return nil
	})
	if err != nil {
		return err
	}
	for _, file := range files {
		fmt.Printf("%s ", file)
	}
	return nil
}
