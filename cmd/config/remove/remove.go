package remove

import (
	"BitSrunLoginGo-CLI/cmd/config/now"
	"BitSrunLoginGo-CLI/cmd/config/use"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/user"
)

func NewConfigRemoveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remove [filename]",
		Aliases: []string{"rm", "del", "delete"},
		Short:   "删除配置",
		Long:    `删除所选的配置`,
		Args:    cobra.RangeArgs(1, 1),
		RunE:    remove,
	}
	return cmd
}

func remove(cmd *cobra.Command, args []string) error {
	name := args[0]

	path, err := user.Current()
	if err != nil {
		return err
	}
	pathdir := path.HomeDir + "/BitSrunLogin/" + name + ".yaml"

	_, err = os.Stat(pathdir)
	if err == nil {
		err = os.Remove(pathdir)
	} else if err != nil {
		return errors.New("cannot find profile:" + name)
	}

	using, err := now.GetNow()
	if err != nil {
		return err
	}
	if using == name {
		fmt.Println("Warning:You are trying to delete a using profile, they may will cause some unexpected case.")
		fmt.Println("now profile: " + name + " is not in used.")
	}
	err = use.Clear()
	if err != nil {
		return err
	}

	fmt.Printf("Removed profile: %s in %s", name, path.HomeDir+"\\BitSrunLogin\\")
	return nil
}
