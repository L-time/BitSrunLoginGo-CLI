package config

import (
	"BitSrunLoginGo-CLI/cmd/config/create"
	"BitSrunLoginGo-CLI/cmd/config/list"
	"BitSrunLoginGo-CLI/cmd/config/now"
	"BitSrunLoginGo-CLI/cmd/config/remove"
	"BitSrunLoginGo-CLI/cmd/config/use"
	"github.com/spf13/cobra"
	"os"
	"os/user"
)

func NewConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config <command>",
		Short: "配置相关参数",
		Long: `为BitSrunLoginGo提供一些需要的配置
所有的配置均是由配置文件持久性存储`,
		PersistentPreRunE: Init,
	}

	cmd.AddCommand(create.NewConfigCreateCmd())
	cmd.AddCommand(remove.NewConfigRemoveCmd())
	cmd.AddCommand(now.NewConfigNowCmd())
	cmd.AddCommand(list.NewConfigListCmd())
	cmd.AddCommand(use.NewConfigUseCmd())

	return cmd
}

// Init 初始化生成一些文件夹
func Init(cmd *cobra.Command, args []string) error {
	path, err := user.Current()
	if err != nil {
		return err
	}
	usedir := path.HomeDir + "/BitSrunLogin/use/"

	//判断文件夹有没有生成
	_, err = os.Stat(path.HomeDir + "/BitSrunLogin")
	if err != nil {
		if os.IsNotExist(err) {
			// 初始化生成文件夹
			err = os.Mkdir(path.HomeDir+"/BitSrunLogin", 0777)
			if err != nil {
				return err
			}
		}
	}

	_, err = os.Stat(usedir)
	if err != nil {
		if os.IsNotExist(err) {
			// 初始化生成文件夹
			err = os.Mkdir(usedir, 0777)
			if err != nil {
				return err
			}
		}
	}
	use, err := os.OpenFile(usedir+"used", os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	use.Close()
	return nil
}
