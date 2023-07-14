package now

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/user"
)

func NewConfigNowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "now",
		Short:   "查看使用的配置文件",
		Long:    "查看当前程序使用的配置文件",
		Aliases: []string{"used"},
		Args:    cobra.NoArgs,
		RunE:    now,
	}
	return cmd
}

func now(cmd *cobra.Command, args []string) error {
	path, err := user.Current()
	if err != nil {
		return err
	}

	used := path.HomeDir + "/BitSrunLogin/use/used"
	using, err := os.OpenFile(used, os.O_RDONLY, 0777)
	if err != nil {
		return err
	}

	usedstr, err := io.ReadAll(using)
	if err != nil {
		return err
	}
	fmt.Println("Now using profile: " + string(usedstr))
	return nil
}

func GetNow() (string, error) {
	path, err := user.Current()
	if err != nil {
		return "", err
	}
	used := path.HomeDir + "/BitSrunLogin/use/used"
	using, err := os.OpenFile(used, os.O_RDONLY, 0777)
	if err != nil {
		return "", err
	}

	usedstr, err := io.ReadAll(using)
	if err != nil {
		return "", err
	}
	return string(usedstr), nil
}
