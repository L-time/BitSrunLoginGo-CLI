package utils

//这里的逻辑是用于未来可能的多配置文件在不同地方的匹配选址问题,目前来说没有被用到是正常的啦233

import (
	"encoding/json"
	"errors"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/user"
)

// users 是config所管理的用户列表
var users map[string]string

func InitUsers(cmd *cobra.Command, args []string) error {
	path, err := user.Current()
	if err != nil {
		return err
	}
	pathdir := path.HomeDir + "/BitSrunLogin/user/" + "user.json"

	_, err = os.Stat(pathdir)
	if err == nil {
		return read()
	}
	_, err = os.Stat(path.HomeDir + "/BitSrunLogin/user")
	if err != nil {
		if os.IsNotExist(err) {
			// 初始化生成文件夹
			err = os.Mkdir(path.HomeDir+"/BitSrunLogin/user", 0777)
			if err != nil {
				return err
			}
		}
	}
	file, err := os.OpenFile(pathdir, os.O_RDWR|os.O_CREATE, 0777)
	defer file.Close()
	if err != nil {
		return err
	}
	return nil
}

// 读取user list
func read() error {
	path, err := user.Current()
	if err != nil {
		return err
	}
	pathdir := path.HomeDir + "/BitSrunLogin/user/" + "user.json"

	file, err := os.OpenFile(pathdir, os.O_RDONLY, 0777)
	defer file.Close()
	if err != nil {
		return err
	}

	jsons, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsons, &users)
	if err != nil {
		return err
	}
	return nil
}

// 写入 user list
func write() error {
	path, err := user.Current()
	if err != nil {
		return err
	}
	pathdir := path.HomeDir + "/BitSrunLogin/user/" + "user.json"

	file, err := os.OpenFile(pathdir, os.O_WRONLY, 0777)
	defer file.Close()
	if err != nil {
		return err
	}

	jsons, err := json.Marshal(&users)
	if err != nil {
		return err
	}

	_, err = io.WriteString(file, string(jsons))
	if err != nil {
		return err
	}
	return nil
}

func List() (map[string]string, error) {
	err := read()
	if err != nil {
		return nil, err
	}
	return users, err

}

func Add(user, path string) error {
	if _, ok := users[user]; ok {
		return errors.New("user already exist")
	}

	users[user] = path
	return write()
}

func Del(user, path string) error {
	delete(users, user)
	return write()
}
