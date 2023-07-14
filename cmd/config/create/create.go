package create

import (
	_ "embed"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/user"
)

//go:embed config.yaml
var yaml string

func NewConfigCreateCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "create [filename]",
		Short:   "初始化配置文件",
		Long:    "生成一份以指定名称的配置文件",
		Aliases: []string{"init", "new", "add"},
		Args:    cobra.RangeArgs(1, 1),
		RunE:    createFile,
	}

	return cmd

}

func createFile(cmd *cobra.Command, args []string) error {
	name := args[0]

	path, err := user.Current()
	if err != nil {
		return err
	}
	pathdir := path.HomeDir + "/BitSrunLogin/" + name + ".yaml"
	//检查是否已经存在了当前配置
	_, err = os.Stat(pathdir)
	if err == nil {
		return errors.New("profile " + name + " already exist")
	}

	file, err := os.OpenFile(pathdir, os.O_RDWR|os.O_CREATE, 0777)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.WriteString(yaml)
	if err != nil {
		return err
	}
	fmt.Printf("Created profile: %s in %s", name, path.HomeDir+"BitSrunLogin\\")
	return nil
}
