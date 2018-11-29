package Api

import (
	"encoding/json"
	"errors"
	"git.docus.tech/kdl12138/DocusServer/Storage"
	"git.docus.tech/kdl12138/DocusServer/Template"
	"io/ioutil"
	"os"
)

func Read(configFile string) (err error) {

	file, err := os.Open(configFile)
	defer file.Close()
	if err != nil {
		return errors.New("配置文件打开错误！请检查文件是否存在以及文件权限是否正确！")
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return errors.New("配置文件解析错误！请参考示例配置文件进行配置！")
	}
	json.Unmarshal([]byte(data), &Template.Setting)
	Server.NewServer(Template.Setting)
	return err
}
