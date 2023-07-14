package use

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/user"
)

func NewConfigUseCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "use [filename]",
		Short:   "使用配置文件",
		Long:    "指定使用的配置文件",
		Aliases: []string{"switch", "change"},
		Args:    cobra.RangeArgs(1, 1),
		RunE:    use,
	}
	return cmd
}

func use(cmd *cobra.Command, args []string) error {
	name := args[0]

	path, err := user.Current()
	if err != nil {
		return err
	}
	pathdir := path.HomeDir + "/BitSrunLogin/" + name + ".yaml"
	//检查是否已经存在了当前配置
	_, err = os.Stat(pathdir)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("profile " + name + " does not exist")
		}
		return err
	}

	//检查是否已经在使用
	used := path.HomeDir + "/BitSrunLogin/use/used"
	using, err := os.OpenFile(used, os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}

	usedByte, err := io.ReadAll(using)
	if err != nil {
		return err
	}
	if name == string(usedByte) {
		return errors.New("this profile has already in used")
	}

	_, err = io.WriteString(using, name)
	if err != nil {
		return err
	}
	fmt.Println("Switch to use profile: " + name)
	return nil
}

func Clear() error {
	path, err := user.Current()
	if err != nil {
		return err
	}
	used := path.HomeDir + "/BitSrunLogin/use/used"
	using, err := os.OpenFile(used, os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	_, err = io.WriteString(using, "")
	if err != nil {
		return err
	}
	return nil
}
